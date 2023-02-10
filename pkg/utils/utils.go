package utils

import (
	"BankCardMS/constant"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
