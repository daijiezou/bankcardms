package worker

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
		WorkerId string `json:"workerId" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Address  string `json:"address"`
		Remarks  string `json:"remarks"`
		Sex      int    `json:"sex" binding:"oneof=1 2"`
	}
	req := new(AddWorkerRequest)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	now := time.Now().UnixMilli()
	worker := &do.Worker{
		WorkerId:   req.WorkerId,
		Name:       req.Name,
		Address:    req.Address,
		Sex:        req.Sex,
		Remarks:    req.Remarks,
		CreateTime: now,
		UpdateTime: now,
	}
	err := mysql.AddWorker(worker)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add worker failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}

func List(c *gin.Context) {
	glog.Info("list workers")
	req := new(utils.CommonListReq)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	result, err := mysql.ListWorkers(req)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add worker failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, result)
	return
}

func Detail(c *gin.Context) {
	workerId := c.Param("worker_id")
	glog.Infof("Detail worker:%v", workerId)
	worker, err := mysql.GetWorker(workerId)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "get worker detail failed"))
		response.Error(c, geminiErr)
		return
	}
	req := mysql.ListBankCardReq{
		WorkerName: worker.Name,
		PageNum:    1,
		PageSize:   99,
	}
	cardList, err := mysql.ListBankCard(&req)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "get worker detail failed"))
		response.Error(c, geminiErr)
		return
	}
	result := new(vo.WorkerDetail)
	result.WorkerInfo = worker
	result.BankCardList = cardList
	var cardIdList []string
	for _, card := range cardList.BankCards {
		cardIdList = append(cardIdList, card.CardId)
	}
	totalIncome, err := gutils.GetTotalIncome(cardIdList)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "calculate total income failed"))
		response.Error(c, geminiErr)
		return
	}
	result.CurrentYearTotalIncome = totalIncome
	response.Success(c, result)
	return
}

func Delete(c *gin.Context) {
	workerId := c.Param("worker_id")
	err := mysql.DeleteWorker(workerId)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add worker failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}

func Update(c *gin.Context) {
	type AddWorkerRequest struct {
		Name    string `json:"name" binding:"required"`
		Address string `json:"address"`
		Sex     int    `json:"sex" binding:"oneof=1 2"`
		Remarks string `json:"remarks"`
	}
	req := new(AddWorkerRequest)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	workerId := c.Param("worker_id")
	worker := &do.Worker{
		Name:       req.Name,
		Address:    req.Address,
		Sex:        req.Sex,
		Remarks:    req.Remarks,
		UpdateTime: time.Now().UnixMilli(),
	}
	err := mysql.UpdateWorker(workerId, worker, "name", "address", "sex", "update_time", "remarks")
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add worker failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}
