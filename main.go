package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-tagle/api"
	"go-tagle/boot"
)

func main() {
	boot.Initialize()
	r := gin.Default()

	api.Register(r)

	err := r.Run(":8000")
	if err != nil {
		fmt.Println("web服务启动失败")
		panic(err)
	}
}
