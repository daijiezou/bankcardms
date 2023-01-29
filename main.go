package main

import (
	"BankCardMS/config"
	"BankCardMS/internal/data/mysql"
	"BankCardMS/internal/pkg/glog"
	"BankCardMS/internal/pkg/middware"
	"BankCardMS/internal/pkg/shutdown"
	"BankCardMS/route/user"
	"BankCardMS/route/worker"
	"context"
	"flag"
	"github.com/gin-gonic/gin"

	"net/http"
)

var (
	configFile = flag.String("config", "", "The path of the configFile")
	env        = flag.String("env", "", "Running Environment")
)

func init() {
	flag.Parse()
	config.ParseConfig("bank-card-ms", *configFile, *env)
	mysql.Init(config.Config.DbDsn)
}

func main() {
	glog.Info("start running")
	r := gin.New()
	r.Use(gin.Recovery(), middware.Cors())
	route(r)
	ctx, cancel := context.WithCancel(context.Background())
	httpSvr := &http.Server{
		Addr:    ":" + config.Config.HttpPort,
		Handler: r,
	}
	go func() {
		if err := httpSvr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Errorf("failed to start http server,error: %v", err)
			return
		}
	}()
	go shutdown.WaitToShutdown(ctx, httpSvr)
	shutdown.HandleSignal(cancel)
}

func route(r *gin.Engine) {
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	user.Router(r)
	worker.Router(r)
}
