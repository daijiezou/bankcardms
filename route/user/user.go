package user

import (
	"BankCardMS/internal/service/user"
	"github.com/gin-gonic/gin"
)

const (
	// RootApi  package_name.version.service
	RootApi = "/gemini/api_server/v1/users"
)

func Router(e *gin.Engine) {
	r := e.Group(RootApi)
	r.POST("/user:action", user.Login)
}
