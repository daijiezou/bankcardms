package bankcard

import (
	"BankCardMS/internal/data/do"
	"BankCardMS/internal/data/mysql"
	"BankCardMS/internal/pkg/gerr"
	"BankCardMS/internal/pkg/glog"
	"BankCardMS/internal/pkg/response"
	"BankCardMS/internal/service/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func Add(c *gin.Context) {
	type AddWorkerRequest struct {
		CardId    string `json:"card_id" binding:"required"`
		CardOwner string `json:"name" binding:"required"`
		BankName  string `json:"bank_name" binding:"required"`
		Remarks   string `json:"remarks"`
	}
	req := new(AddWorkerRequest)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	now := time.Now().UnixMilli()
	bankCard := &do.BankCard{
		CardId:     req.CardId,
		CardOwner:  req.CardOwner,
		BankName:   req.BankName,
		Remarks:    req.Remarks,
		CreateTime: now,
		UpdateTime: now,
	}
	err := mysql.AddBankCard(bankCard)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add bankCard failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}

func List(c *gin.Context) {
	req := new(utils.CommonListReq)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	result, err := mysql.ListBankCard(req)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add worker failed"))
		response.Error(c, geminiErr)
		return
	}
	//todo 当年总收入
	response.Success(c, result)
	return
}

func Detail(c *gin.Context) {
	cardId := c.Param("card_id")
	bankCard, err := mysql.GetBankCard(cardId)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add bankCard failed"))
		response.Error(c, geminiErr)
		return
	}
	//todo 近三年总收入
	response.Success(c, bankCard)
	return
}

func Delete(c *gin.Context) {
	cardId := c.Param("card_id")
	err := mysql.DeleteBankCard(cardId)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add bankCard failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}

func Update(c *gin.Context) {
	type AddWorkerRequest struct {
		CardOwner string `json:"name" binding:"required"`
		BankName  string `json:"bank_name" binding:"required"`
		Remarks   string `json:"remarks"`
	}
	req := new(AddWorkerRequest)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	cardId := c.Param("card_id")
	now := time.Now().UnixMilli()
	bankCard := &do.BankCard{
		CardOwner:  req.CardOwner,
		BankName:   req.BankName,
		Remarks:    req.Remarks,
		CreateTime: now,
		UpdateTime: now,
	}
	err := mysql.UpdateBankCard(cardId, bankCard, "card_owner", "bank_name", "remarks", "update_time")
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add worker failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}
