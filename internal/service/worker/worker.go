package worker

import (
	"BankCardMS/internal/data/do"
	"BankCardMS/internal/data/mysql"
	"BankCardMS/internal/pkg/gerr"
	"BankCardMS/internal/pkg/glog"
	"BankCardMS/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

func Add(c *gin.Context) {
	type AddWorkerRequest struct {
		WorkerId string `json:"workerId" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Address  string `json:"address"`
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
	req := new(do.WorkerListReq)
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
	worker, err := mysql.GetWorker(workerId)
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add worker failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, worker)
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
		UpdateTime: time.Now().UnixMilli(),
	}
	err := mysql.UpdateWorker(workerId, worker, "name", "address", "sex", "update_time")
	if err != nil {
		geminiErr := gerr.FromError(err)
		glog.Errorf(geminiErr.ErrorWithMsg(err, "add worker failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}
