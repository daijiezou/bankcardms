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
		WorkerId string `json:"workerId"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		Sex      int    `json:"sex"`
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
		glog.Errorf(geminiErr.ErrorWithMsg(c, err, "add worker failed"))
		response.Error(c, geminiErr)
		return
	}
	response.Success(c, nil)
	return
}

func List(c *gin.Context) {

}

func Detail(c *gin.Context) {

}
