package gerr

type ResCode int64

const (
	/*
		0 请求成功
		错误码定义规则
		6位 2+1+3 application+family+bizCode
		family：  3 流程处理相关 4 客户端错误 5 系统服务
	*/
	CodeSuccess ResCode = 0

	ErrCodeUnauthorized     ResCode = 154001
	ErrCodeDataNotFound     ResCode = 154004
	ErrCodePermissionDenied ResCode = 154006
	ErrCodeLoginError       ResCode = 154012
	ErrCodeLoginAuthError   ResCode = 154013
	ErrCodeWrongParam       ResCode = 154014
	ErrCodeInvalidPassword  ResCode = 154015

	// 0-29 warning 30+ error
	ErrCodeServerBusy   ResCode = 155030
	ErrCodeDbError      ResCode = 155031
	ErrCodeNetworkError ResCode = 155032
)

var errorText = map[ResCode]string{
	CodeSuccess: "success",

	ErrCodeUnauthorized:     "认证错误:无效的Token",
	ErrCodeDataNotFound:     "数据错误:数据不存在",
	ErrCodePermissionDenied: "认证错误:权限不足",
	ErrCodeLoginError:       "登录失败:用户名或密码不正确",
	ErrCodeWrongParam:       "数据错误:请求参数不正确",
	ErrCodeLoginAuthError:   "认证错误:用户未授权",

	ErrCodeServerBusy:      "系统错误:服务繁忙，请稍后再试",
	ErrCodeDbError:         "系统错误:数据库错误",
	ErrCodeNetworkError:    "系统错误:网络错误",
	ErrCodeInvalidPassword: "用户名或密码错误",
}

func (r ResCode) Msg() string {
	msg, ok := errorText[r]
	if !ok {
		return errorText[ErrCodeServerBusy]
	}
	return msg
}

func (r ResCode) Level() string {
	if (r/100)%100 == 50 && (r%100 >= 30) {
		return "ERROR"
	} else {
		return "WARNING"
	}
}
