package mysql

import (
	"BankCardMS/internal/pkg/glog"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var apiClient *xorm.Engine

func Init(dbDsn string) {
	var err error
	apiClient, err = xorm.NewEngine("mysql", dbDsn)
	if err != nil {
		glog.Fatalf("init mysql err:", err.Error())
	}
	apiClient.ShowSQL(true)
	apiClient.SetMaxIdleConns(20)

	apiClient.Logger().SetLevel(log.LOG_WARNING)
	return
}

func MySQL() *xorm.Engine {
	return apiClient
}
