package mysql

import (
	"BankCardMS/data/do"
	"BankCardMS/pkg/gerr"
	"BankCardMS/pkg/glog"
	"github.com/pkg/errors"
)

func AddWorker(worker *do.Worker) error {
	glog.Infof("worker:%+v", worker)
	_, err := MySQL().Insert(worker)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return errors.WithStack(gErr)
	}
	return nil
}

func DeleteWorker(workerId string) error {
	worker := new(do.Worker)
	count, err := MySQL().Where("worker_id = ?", workerId).Delete(worker)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return errors.WithStack(gErr)
	}
	if count == 0 {
		gErr := &gerr.GeminiError{
			Code: gerr.ErrCodeDataNotFound,
			Err:  gerr.ErrDataNotFound,
		}
		return errors.WithStack(gErr)
	}
	return nil
}

func GetWorker(workerId string) (*do.Worker, error) {
	worker := new(do.Worker)
	has, err := MySQL().Table("worker").Where("worker_id = ?", workerId).Get(worker)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return worker, errors.WithStack(gErr)
	}
	if !has {
		gErr := &gerr.GeminiError{
			Code: gerr.ErrCodeDataNotFound,
			Err:  gerr.ErrDataNotFound,
		}
		return worker, errors.WithStack(gErr)
	}
	return worker, nil
}

func UpdateWorker(workerId string, worker *do.Worker, cols ...string) error {
	_, err := MySQL().Table("worker").Where("worker_id = ?", workerId).Cols(cols...).Update(worker)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return errors.WithStack(gErr)
	}
	return nil
}

type WorkerReq struct {
	WorkerId   string `json:"worker_id" form:"worker_id"`
	WorkerName string `json:"worker_name" form:"worker_name"`
	PageNum    int    `json:"page_num" form:"page_num" binding:"min=1"`
	PageSize   int    `json:"page_size" form:"page_size" binding:"min=1,max=2000"`
}

func ListWorkers(req *WorkerReq) (result *do.WorkerList, err error) {
	result = new(do.WorkerList)
	session := MySQL().Table("worker").Select("*").And("delete_time = ?", 0)
	if req.WorkerName != "" {
		session.And("name like ?", "%"+req.WorkerName+"%")

	}
	if req.WorkerId != "" {
		session.And("worker_id like ?", "%"+req.WorkerId+"%")
	}

	count, err := session.
		Limit(req.PageSize, req.PageSize*(req.PageNum-1)).
		Desc("worker.create_time").
		FindAndCount(&result.Workers)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return nil, errors.Wrapf(gErr, "req:%v,", req)
	}
	result.ListCount = int(count)
	return
}
