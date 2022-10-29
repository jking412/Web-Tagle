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
	//r.Use(middleware.Cors())

	r.Any("/ping", ping)

	uc := new(controller.UserController)
	hc := new(controller.HomeController)

	r.GET("/", hc.Index)

	userGroup := r.Group("/user")
	{
		userGroup.GET("/register", uc.Register)
		userGroup.POST("/register", uc.DoRegister)
		userGroup.GET("/login", uc.Login)
		userGroup.POST("/login", uc.DoLogin)

		githubGroup := userGroup.Group("/github")
		{
			githubGroup.GET("/login", uc.GithubLogin)
			githubGroup.GET("/oauth2", uc.GithubLoginCallback)
		}

		emailGroup := userGroup.Group("/email")
		{
			emailGroup.GET("/activate", uc.ActivateEmail)
			emailGroup.GET("/send", uc.SendEmailVerifyCode)
			emailGroup.GET("/login", uc.EmailLogin)
			emailGroup.POST("/login", uc.DoEmailLogin)
		}
	}

	r.Use(middleware.Auth())

	habitGroup := r.Group("/habit")
	{
		habitGroup.GET("/all", controller.GetAllHabits)
		habitGroup.POST("/create", controller.CreateHabit)
		habitGroup.POST("/update", controller.UpdateHabit)
		habitGroup.POST("/finish", controller.UpdateHabitFinishedTime)
		habitGroup.POST("/unfinish", controller.UpdateHabitUnfinishedTime)
		habitGroup.POST("/delete", controller.DeleteHabit)
	}

	taskGroup := r.Group("/task")
	{
		taskGroup.GET("/all", controller.GetAllTasks)
		taskGroup.POST("/create", controller.CreateTask)
		taskGroup.POST("/update", controller.UpdateTask)
		taskGroup.POST("/finish", controller.UpdateFinishedTime)
		taskGroup.POST("/delete", controller.DeleteTask)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}
