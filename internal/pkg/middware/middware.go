package middware

import (
	"BankCardMS/internal/constant"
	"BankCardMS/internal/pkg/gerr"
	"BankCardMS/internal/pkg/glog"
	"BankCardMS/internal/pkg/jwt"
	"BankCardMS/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	UserAction  = "user"
	AdminAction = "admin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,spaceid")
		c.Header("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS,UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type,spaceid")
		c.Header("Access-Control-Request-Headers", "spaceid")
		c.Header("Access-Control-Request-Headers", "pagename")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.UnauthorizedError(c, gerr.ErrCodeUnauthorized)
			glog.Errorf("token not found: %v", authHeader)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			glog.Errorf("invalid token")
			response.UnauthorizedError(c, gerr.ErrCodeUnauthorized)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		claim, err := jwt.ParseToken(parts[1])
		if err != nil {
			glog.Errorf("invalid token,err:%v", err)
			response.UnauthorizedError(c, gerr.ErrCodeUnauthorized)
			c.Abort()
			return
		}
		c.Set(constant.CtxUserNameKey, claim.UserName)
		c.Set(constant.CtxUserIdKey, claim.UserId)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
