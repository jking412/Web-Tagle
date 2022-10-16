package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")
		if userId == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}
