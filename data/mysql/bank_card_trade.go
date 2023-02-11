package mysql

import (
	"BankCardMS/data/do"
	"BankCardMS/pkg/gerr"
	"BankCardMS/service/utils"
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

func ListBankCardTrade(req *utils.CommonListReq, bankCardId string) (result *do.BankCardTradeList, err error) {
	result = new(do.BankCardTradeList)
	session := MySQL().Table("bank_card_trade").Select("*").And("delete_time = ?", 0).
		Where("card_id = ?", bankCardId)
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

func GetBankCardIncome(bankCardId string, start, end int64) (result int64, err error) {
	count, err := MySQL().Table("bank_card_trade").Where("card_id = ?", bankCardId).
		And("delete_time = 0").
		Where("trade_time >= ? and trade_time <= ?", start, end).Sum(new(do.BankCardTrade), "trade_amount")
	if err != nil {
		gErr := &gerr.GeminiError{
			Code:     gerr.ErrCodeDbError,
			CauseMsg: err.Error(),
		}
		return int64(count), errors.Wrapf(gErr, "bankCardId:%v,", bankCardId)
	}
	return int64(count), nil
}
