package api

import (
	"github.com/gin-gonic/gin"
	"go-tagle/app/core"
)

func (ua *UserApi) SendCaptcha(c *gin.Context) {
	data, statusCode := core.GenerateCaptcha(c)
	c.JSON(statusCode, data)
}

func (ua *UserApi) SendEmailCaptcha(c *gin.Context) {
	data, statusCode := core.SendEmailCaptcha(c)
	c.JSON(statusCode, data)
}

func (ua *UserApi) EmailLogin(c *gin.Context) {
	data, statusCode := core.EmailLogin(c)
	c.JSON(statusCode, data)
}
