package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-tagle/app/api"
	controller2 "go-tagle/app/controller"
	"go-tagle/middleware"
	"go-tagle/pkg/session"
)

func Register(r *gin.Engine) {
	r.Use(sessions.Sessions("mysession", session.Store))
	//r.Use(middleware.Cors())

	r.Any("/ping", ping)

	uc := new(controller2.UserController)
	hc := new(controller2.HomeController)

	ua := new(api.UserApi)

	r.GET("/", hc.Index)

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/ping", ping)
		userApiGroup := apiGroup.Group("/user")
		{
			userApiGroup.POST("/login", ua.Login)
			userApiGroup.POST("/register", ua.Register)

			emailGroup := apiGroup.Group("/email")
			{
				emailGroup.GET("/captcha", ua.SendCaptcha)
				emailGroup.POST("/send", ua.SendEmailCaptcha)
				emailGroup.POST("/login", ua.EmailLogin)
			}

		}

	}

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
		habitGroup.GET("/all", controller2.GetAllHabits)
		habitGroup.POST("/create", controller2.CreateHabit)
		habitGroup.POST("/update", controller2.UpdateHabit)
		habitGroup.POST("/finish", controller2.UpdateHabitFinishedTime)
		habitGroup.POST("/unfinish", controller2.UpdateHabitUnfinishedTime)
		habitGroup.POST("/delete", controller2.DeleteHabit)
	}

	taskGroup := r.Group("/task")
	{
		taskGroup.GET("/all", controller2.GetAllTasks)
		taskGroup.POST("/create", controller2.CreateTask)
		taskGroup.POST("/update", controller2.UpdateTask)
		taskGroup.POST("/finish", controller2.UpdateFinishedTime)
		taskGroup.POST("/delete", controller2.DeleteTask)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}
