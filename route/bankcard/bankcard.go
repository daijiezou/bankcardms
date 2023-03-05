package bankcard

import (
	"BankCardMS/pkg/middware"
	"BankCardMS/service/bankcard"
	"github.com/gin-gonic/gin"
)

const (
	// RootApi  package_name.version.service
	RootApi = "/bank_card_ms/api_server/v1/bank_cards"
)

func Router(e *gin.Engine) {
	r := e.Group(RootApi)
	r.Use(middware.JWTAuthMiddleware())
	r.POST("/bank_card", bankcard.Add)               //添加银行卡
	r.GET("", bankcard.List)                         //查看银行卡列表
	r.GET("/bank_card/:card_id", bankcard.Detail)    //查看银行卡详情
	r.DELETE("/bank_card/:card_id", bankcard.Delete) //删除银行卡
	r.PUT("/bank_card/:card_id", bankcard.Update)    //更新银行卡信息

	r.POST("/bank_card/:card_id/trades", bankcard.AddTrade) //添加银行卡交易记录
	//r.GET("/bank_card/:card_id/trades", bankcard.ListTrade)                //查看银行卡交易流水
	r.DELETE("/bank_card/:card_id/trades/:trade_id", bankcard.DeleteTrade) //删除银行卡交易流水
	r.PUT("/bank_card/:card_id/trades/:trade_id", bankcard.UpdateTrade)    //更新银行卡交易流水
}
