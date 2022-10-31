package controller

import (
	"github.com/gin-gonic/gin"
	"go-tagle/app/core"
	"go-tagle/model/user"
	"go-tagle/pkg/redislib"
	"go-tagle/pkg/session"
	"net/http"
)

func (uc *UserController) EmailLogin(c *gin.Context) {
	_, ok := session.GetUser(c)
	if ok {
		c.Redirect(http.StatusFound, "/")
		return
	}
	data, statusCode := core.GenerateCaptcha(c)
	c.HTML(statusCode, "login-email", data)
}

func (uc *UserController) SendEmailVerifyCode(c *gin.Context) {
	data, statusCode := core.SendEmailCaptcha(c)
	c.HTML(statusCode, "login-email", data)
}

func (uc *UserController) DoEmailLogin(c *gin.Context) {
	data, statusCode := core.EmailLogin(c)
	if statusCode == http.StatusOK {
		c.Redirect(http.StatusFound, "/")
	} else {
		c.HTML(statusCode, "error", data)
	}
}

func (uc *UserController) ActivateEmail(c *gin.Context) {
	_email := c.Query("email")
	code := c.Query("code")
	if code == "" {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "链接无效",
		})
		return
	}
	redisCode := redislib.GlobalRedis.Get(_email)
	if redisCode == "" || redisCode != code {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "链接无效",
		})
		return
	}
	if ok := user.ActivateEmail(_email); !ok {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"msg": "激活失败",
		})
		return
	}
	redislib.GlobalRedis.Del(_email)
	c.Redirect(http.StatusFound, "/")
}
