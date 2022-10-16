package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-tagle/controller"
	"go-tagle/middleware"
	"go-tagle/pkg/session"
)

func Register(r *gin.Engine) {
	r.Use(sessions.Sessions("mysession", session.Store))
	r.Use(middleware.Cors())

	r.Any("/ping", ping)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controller.Register)
		userGroup.POST("/login", controller.Login)
	}

	r.Use(middleware.Auth())
	habitGroup := r.Group("/habit")
	{
		habitGroup.POST("/create", controller.CreateHabit)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}
