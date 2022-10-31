package api

import (
	"github.com/gin-gonic/gin"
	"go-tagle/app/core"
	"net/http"
)

type UserApi struct {
}

func (ua *UserApi) Login(c *gin.Context) {
	msg, statusCode := core.Login(c)
	if statusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	} else {
		c.JSON(statusCode, gin.H{
			"msg": msg,
		})
	}
}

func (ua *UserApi) Register(c *gin.Context) {
	msg, statusCode := core.Register(c)
	if statusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	} else {
		c.JSON(statusCode, gin.H{
			"msg": msg,
		})
	}
}
