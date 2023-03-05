package worker

import (
	"BankCardMS/pkg/middware"
	"BankCardMS/service/worker"
	"github.com/gin-gonic/gin"
)

const (
	// RootApi  package_name.version.service
	RootApi = "/bank_card_ms/api_server/v1/workers"
)

func Router(e *gin.Engine) {
	r := e.Group(RootApi)
	r.Use(middware.JWTAuthMiddleware())
	r.POST("/worker", worker.Add)
	r.GET("", worker.List)
	r.GET("/worker/:worker_id", worker.Detail)
	r.DELETE("/worker/:worker_id", worker.Delete)
	r.PUT("/worker/:worker_id", worker.Update)
}
