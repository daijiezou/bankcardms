package shutdown

import (
	"context"
	"github.com/golang/glog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	wg sync.WaitGroup
)

func HandleSignal(cancel context.CancelFunc) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-signalCh:
			switch s {
			case syscall.SIGINT:
				fallthrough
			case syscall.SIGTERM:
				fallthrough
			default:
				glog.Info("stop signal received, stop all server")
				cancel()
				wg.Wait()
				glog.Info("gracefully exit")
				return
			}
		}
	}
}

func WaitToShutdown(ctx context.Context, svr *http.Server) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			shutdown(svr)
			return
		}
	}
}

func shutdown(svr *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		glog.Errorf("failed to shut down server:%v", err)
		return
	}
	glog.Info("shutdown server success")
}
