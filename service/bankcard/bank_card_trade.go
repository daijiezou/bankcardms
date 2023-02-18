package bankcard

import (
	"BankCardMS/data/do"
	"BankCardMS/data/mysql"
	"BankCardMS/pkg/gerr"
	"BankCardMS/pkg/glog"
	"BankCardMS/pkg/idgenerator"
	"BankCardMS/pkg/response"
	"BankCardMS/service/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func AddTrade(c *gin.Context) {
	cardId := c.Param("card_id")
	type AddTradeReq struct {
		TradeAmount int    `json:"trade_amount" binding:"required,min=1"`
		Remarks     string `json:"remarks" binding:"required"`
		TradeTime   int64  `json:"trade_time" binding:"required"`
	}
	req := new(AddTradeReq)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}

	_, err := mysql.GetBankCard(cardId)
	if err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCodeWithMsg(c, gerr.ErrCodeWrongParam, "银行卡号不存在")
		return
	}

	now := time.Now().UnixMilli()
	trade := &do.BankCardTrade{
		TradeId:     idgenerator.GenIDString(),
		CardId:      cardId,
		TradeAmount: req.TradeAmount,
		Remarks:     req.Remarks,
		TradeTime:   req.TradeTime,
		CreateTime:  now,
		UpdateTime:  now,
		DeleteTime:  0,
	}
	err = mysql.AddBankCardTrade(trade)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add bankCard trade failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, gin.H{
		"trade_id": trade.TradeId,
	})
	return
}

func ListTrade(c *gin.Context) {
	cardId := c.Param("card_id")
	req := new(utils.CommonListReq)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	result, err := mysql.ListBankCardTrade(req, cardId)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add worker failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, result)
	return
}

func DeleteTrade(c *gin.Context) {
	tradeId := c.Param("trade_id")
	err := mysql.DeleteBankCardTrade(tradeId)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "delete trade failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}

func UpdateTrade(c *gin.Context) {
	tradeId := c.Param("trade_id")
	type AddTradeReq struct {
		TradeAmount int    `json:"trade_amount" binding:"required,min=1"`
		Remarks     string `json:"remarks" binding:"required"`
		TradeTime   int64  `json:"trade_time" binding:"required"`
	}
	req := new(AddTradeReq)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	now := time.Now().UnixMilli()
	trade := &do.BankCardTrade{
		TradeAmount: req.TradeAmount,
		Remarks:     req.Remarks,
		UpdateTime:  now,
		TradeTime:   req.TradeTime,
	}
	err := mysql.UpdateBankCardTrade(tradeId, trade, "trade_amount", "remarks", "update_time", "trade_time")
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "update trade failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}
