package response

import (
	"BankCardMS/internal/pkg/gerr"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, err *gerr.GeminiError) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: err.Code,
		Msg:  err.Error(),
	})
}

func ErrorCode(c *gin.Context, errCode gerr.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: errCode,
		Msg:  errCode.Msg(),
	})
}

func ErrorCodeWithMsg(c *gin.Context, errCode gerr.ResCode, msg string) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: errCode,
		Msg:  msg,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: 0,
		Msg:  gerr.CodeSuccess.Msg(),
		Data: data,
	})
}

func UnauthorizedError(c *gin.Context, errCode gerr.ResCode) {
	c.JSON(http.StatusUnauthorized, &ResponseData{
		Code: errCode,
		Msg:  errCode.Msg(),
	})
}

type ResponseData struct {
	Code gerr.ResCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data,omitempty" `
}
