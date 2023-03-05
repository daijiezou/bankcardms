package utils

import (
	"BankCardMS/constant"
	"BankCardMS/data/mysql"
	"BankCardMS/pkg/gerr"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
)

func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	var ErrorUserNotLogin = errors.New("用户未登录")
	uid, ok := c.Get(constant.CtxUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func GetTotalIncome(cardIdList []string) (totalIncome int64, err error) {
	totalIncome = 0
	for _, cardId := range cardIdList {
		startTime, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%d-01-01 00:00:00", time.Now().Year()), time.Local)
		if err != nil {
			gErr := &gerr.GeminiError{
				Code:     gerr.ErrCodeServerBusy,
				CauseMsg: err.Error(),
			}
			return totalIncome, errors.WithStack(gErr)
		}
		endTime, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%d-12-31 23:59:59", time.Now().Year()), time.Local)
		if err != nil {
			gErr := &gerr.GeminiError{
				Code:     gerr.ErrCodeServerBusy,
				CauseMsg: err.Error(),
			}
			return totalIncome, errors.WithStack(gErr)
		}

		singleCardIncome, err := mysql.GetBankCardIncome(cardId, startTime.UnixMilli(), endTime.UnixMilli())
		if err != nil {
			return 0, err
		}
		totalIncome += singleCardIncome
	}
	return totalIncome, nil
}
