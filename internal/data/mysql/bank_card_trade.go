package mysql

import (
	"BankCardMS/internal/data/do"
	"BankCardMS/internal/pkg/gerr"
	"BankCardMS/internal/service/commonreq"
	"BankCardMS/internal/service/vo"
	"github.com/pkg/errors"
)

func AddBankCardTrade(bankCardTrade *do.BankCardTrade) error {
	_, err := MySQL().Insert(bankCardTrade)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return errors.WithStack(gErr)
	}
	return nil
}

func DeleteBankCardTrade(tradeId string) error {
	bankCardTrade := new(do.BankCardTrade)
	count, err := MySQL().Where("trade_id = ?", tradeId).Delete(bankCardTrade)
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

func GetBankCardTrade(tradeId string) (*do.BankCardTrade, error) {
	bankCardTrade := new(do.BankCardTrade)
	has, err := MySQL().Where("trade_id = ?", tradeId).Get(bankCardTrade)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return bankCardTrade, errors.WithStack(gErr)
	}
	if !has {
		gErr := &gerr.GeminiError{
			Code: gerr.ErrCodeDataNotFound,
			Err:  gerr.ErrDataNotFound,
		}
		return bankCardTrade, errors.WithStack(gErr)
	}
	return bankCardTrade, nil
}

func UpdateBankCardTrade(tradeId string, bankCardTrade *do.BankCardTrade, cols ...string) error {
	_, err := MySQL().Where("trade_id = ?", tradeId).Cols(cols...).Update(bankCardTrade)
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return errors.WithStack(gErr)
	}
	return nil
}

func ListBankCardTrade(req *commonreq.CommonListReq) (result *vo.BankCardTradeList, err error) {
	result = new(vo.BankCardTradeList)
	session := MySQL().Select("*").And("delete_time = ?", 0)
	if req.Filter != "" {
		session.And("name like ?", "%"+req.Filter+"%")

	}
	count, err := session.
		Limit(req.PageSize, req.PageSize*(req.PageNum-1)).
		Desc("create_time").
		FindAndCount(&result.BankCardTrades)
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
