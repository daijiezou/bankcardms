package middware

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
