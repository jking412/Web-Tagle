package main

import (
	"github.com/gin-gonic/gin"
	"go-tagle/api"
	"go-tagle/boot"
	"go-tagle/conf"
	"go-tagle/pkg/logger"
	"go-tagle/pkg/template"
)

func main() {
	boot.Initialize()

	r := gin.Default()

	template.InitTemplate(r)

	api.Register(r)

	err := r.Run(":" + conf.ServerConf.Port)
	if err != nil {
		logger.ErrorString("web", "启动失败", err.Error())
		panic(err)
	}
}
