package main

import (
	"github.com/gin-gonic/gin"
	"go-tagle/api"
	"go-tagle/boot"
	"go-tagle/pkg/logger"
	"go-tagle/pkg/viperlib"
)

func main() {
	boot.Initialize()

	r := gin.Default()

	api.Register(r)

	logger.ErrorString("test", "test", "test")

	err := r.Run(":" + viperlib.GetString("server.port"))
	if err != nil {
		logger.ErrorString("web", "启动失败", err.Error())
		panic(err)
	}
}
