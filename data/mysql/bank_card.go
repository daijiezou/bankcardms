package mysql

import (
	"BankCardMS/data/do"
	"BankCardMS/pkg/gerr"
	"BankCardMS/service/utils"
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

func GetBankCard(cardId string) (*do.BankCard, error) {
	bankCard := new(do.BankCard)
	has, err := MySQL().Where("card_id = ?", cardId).Get(bankCard)
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

func ListBankCard(req *utils.CommonListReq) (result *do.BankCardList, err error) {
	result = new(do.BankCardList)
	session := MySQL().Table("worker").Select("*").And("delete_time = ?", 0)
	if req.Filter != "" {
		session.And("name like ?", "%"+req.Filter+"%")

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
