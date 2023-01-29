package mysql

import (
	"BankCardMS/internal/data/do"
	"BankCardMS/internal/pkg/gerr"
	"github.com/pkg/errors"
)

func AddWorker(worker *do.Worker) error {
	_, err := MySQL().Table("worker").Insert(worker)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return errors.WithStack(gErr)
	}
	return nil
}
