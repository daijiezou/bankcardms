package bankcard

import (
	"BankCardMS/data/do"
	"BankCardMS/data/mysql"
	"BankCardMS/pkg/gerr"
	"BankCardMS/pkg/glog"
	"BankCardMS/pkg/response"
	gutils "BankCardMS/pkg/utils"
	"BankCardMS/service/utils"
	"BankCardMS/service/vo"
	"github.com/gin-gonic/gin"
	"time"
)

func Add(c *gin.Context) {
	type AddWorkerRequest struct {
		CardId    string `json:"card_id" binding:"required"`
		CardOwner string `json:"card_owner" binding:"required"`
		BankName  string `json:"bank_name" binding:"required"`
		Remarks   string `json:"remarks"`
	}
	req := new(AddWorkerRequest)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	// 校验worker是否存在
	_, err := mysql.GetWorker(req.CardOwner)
	if err != nil {
		response.ErrorCodeWithMsg(c, gerr.ErrCodeWrongParam, "银行卡所有者不存在")
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
	err = mysql.AddBankCard(bankCard)
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
	req := new(mysql.ListBankCardReq)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	glog.Infof("req:%+v", req)
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
	req := new(utils.CommonListReq)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	cardId := c.Param("card_id")
	bankCard, err := mysql.GetBankCard(cardId)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "get bankCard detail failed"))
		response.Error(c, geminiErr)
		return
	}
	tradeList, err := mysql.ListBankCardTrade(req, cardId)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "list bankCard trade failed"))
		response.Error(c, geminiErr)
		return
	}
	totalIncome, err := gutils.GetTotalIncome([]string{cardId})
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "list bankCard trade failed"))
		response.Error(c, geminiErr)
		return
	}
	result := vo.BankCardDetail{
		BankCardInfo:           bankCard,
		TradeList:              tradeList,
		CurrentYearTotalIncome: totalIncome,
	}
	response.Success(c, result)
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
		CardOwner string `json:"card_owner" binding:"required"`
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
	_, err := mysql.GetBankCard(cardId)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "get bankCard detail failed"))
		response.Error(c, geminiErr)
		return
	}

	// 校验worker是否存在
	_, err = mysql.GetWorker(req.CardOwner)
	if err != nil {
		response.ErrorCodeWithMsg(c, gerr.ErrCodeWrongParam, "银行卡所有者不存在")
		return
	}

	now := time.Now().UnixMilli()
	bankCard := &do.BankCard{
		CardOwner:  req.CardOwner,
		BankName:   req.BankName,
		Remarks:    req.Remarks,
		CreateTime: now,
		UpdateTime: now,
	}
	err = mysql.UpdateBankCard(cardId, bankCard, "card_owner", "bank_name", "remarks", "update_time")
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "update bank card info failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}
