package user

import (
	"BankCardMS/internal/pkg/middware"
	"BankCardMS/internal/service/user"
	"github.com/gin-gonic/gin"
)

const (
	// RootApi  package_name.version.service
	RootApi = "/bank_card_ms/api_server/v1/users"
)

func Router(e *gin.Engine) {
	r := e.Group(RootApi)
	r.POST("/user/login", user.Login)
	r.POST("/user/editpwd", middware.JWTAuthMiddleware(), user.EditPwd)
}
