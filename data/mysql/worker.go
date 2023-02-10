package mysql

import (
	"BankCardMS/data/do"
	"BankCardMS/pkg/gerr"
	"BankCardMS/pkg/glog"
	"BankCardMS/service/utils"
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

func ListWorkers(req *utils.CommonListReq) (result *do.WorkerList, err error) {
	result = new(do.WorkerList)
	session := MySQL().Table("worker").Select("*").And("delete_time = ?", 0)
	if req.Filter != "" {
		session.And("name like ?", "%"+req.Filter+"%")

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
