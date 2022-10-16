package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-tagle/model"
	"net/http"
)

type CreateHabitReq struct {
	Name string `json:"name" valid:"name"`
}

func CreateHabit(c *gin.Context) {
	createHabitReq := &CreateHabitReq{}
	if err := c.ShouldBindJSON(&createHabitReq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}

	session := sessions.Default(c)
	userId := session.Get("userId").(int)
	habit := &model.Habit{
		Name:   createHabitReq.Name,
		UserId: userId,
	}
	if err := habit.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "创建成功",
		"habit": habit,
	})
}

func UpdateHabit(c *gin.Context) {
}
