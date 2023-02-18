package mysql

import (
	"BankCardMS/data/do"
	"BankCardMS/pkg/gerr"
	"github.com/pkg/errors"
)

func AddBankCard(bankCard *do.BankCard) error {
	_, err := MySQL().Insert(bankCard)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return errors.WithStack(gErr)
	}
	return nil
}

func DeleteBankCard(cardId string) error {
	bankCard := new(do.BankCard)
	count, err := MySQL().Where("card_id = ?", cardId).Delete(bankCard)
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

func GetBankCard(cardId string) (*do.BankCardDetail, error) {
	bankCard := new(do.BankCardDetail)
	has, err := MySQL().Table("bank_card").Select("bank_card.*,worker.name").
		Where("card_id = ?", cardId).Join("INNER", "worker", "bank_card.card_owner = worker.worker_id").Get(bankCard)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return bankCard, errors.WithStack(gErr)
	}
	if !has {
		gErr := &gerr.GeminiError{
			Code: gerr.ErrCodeDataNotFound,
			Err:  gerr.ErrDataNotFound,
		}
		return bankCard, errors.WithStack(gErr)
	}
	return bankCard, nil
}

func UpdateBankCard(cardId string, bankCard *do.BankCard, cols ...string) error {
	_, err := MySQL().Where("card_id = ?", cardId).Cols(cols...).Update(bankCard)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return errors.WithStack(gErr)
	}
	return nil
}

type ListBankCardReq struct {
	BankName   string `json:"bank_name" form:"bank_name"`
	WorkerName string `json:"worker_name" form:"worker_name"`
	PageNum    int    `json:"page_num" form:"page_num" binding:"min=1"`
	PageSize   int    `json:"page_size" form:"page_size" binding:"min=1,max=100"`
}

func ListBankCard(req *ListBankCardReq) (result *do.BankCardList, err error) {
	result = new(do.BankCardList)
	session := MySQL().Table("bank_card").Select("bank_card.*,worker.name").And("bank_card.delete_time = ?", 0).
		Join("INNER", "worker", "bank_card.card_owner = worker.worker_id")
	if req.BankName != "" {
		session.And("bank_card.bank_name like ?", "%"+req.BankName+"%")
	}
	if req.WorkerName != "" {
		session.And("worker.name like ?", "%"+req.WorkerName+"%")
	}
	count, err := session.
		Limit(req.PageSize, req.PageSize*(req.PageNum-1)).
		Desc("create_time").
		FindAndCount(&result.BankCards)
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
