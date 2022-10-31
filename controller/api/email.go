package api

import (
	"github.com/gin-gonic/gin"
	"go-tagle/controller/core"
)

func (ua *UserApi) SendCaptcha(c *gin.Context) {
	data, statusCode := core.GenerateCaptcha(c)
	c.JSON(statusCode, data)
}
