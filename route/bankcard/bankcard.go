package bankcard

import (
	"BankCardMS/internal/pkg/middware"
	"BankCardMS/internal/service/bankcard"
	"github.com/gin-gonic/gin"
)

const (
	// RootApi  package_name.version.service
	RootApi = "/bank_card_ms/api_server/v1/bank_cards"
)

func Router(e *gin.Engine) {
	r := e.Group(RootApi)
	r.Use(middware.JWTAuthMiddleware())
	r.POST("", bankcard.Add)
	r.GET("", bankcard.List)
	r.GET("/bank_card/:card_id", bankcard.Detail)
	r.DELETE("/bank_card/:card_id", bankcard.Delete)
	r.PUT("/bank_card/:card_id", bankcard.Update)
}
