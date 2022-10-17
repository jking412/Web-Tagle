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

type UpdateHabitReq struct {
	Id             int    `json:"id,omitempty" valid:"id"`
	Name           string `json:"name" valid:"name"`
	Desc           string `json:"desc" valid:"desc"`
	Difficulty     int    `json:"difficulty" valid:"difficulty"`
	Tag            string `json:"tag" valid:"tag"`
	UserId         int    `json:"userId" valid:"userId"`
	FinishedTime   int    `json:"finishedTime" valid:"finishedTime"`
	UnFinishedTime int    `json:"unfinishedTime" valid:"unfinishedTime"`
}

type UpdateHabitFinishedTimeReq struct {
	Id           int `json:"id,omitempty" valid:"id"`
	FinishedTime int `json:"finishedTime" valid:"finishedTime"`
}

type UpdateHabitUnfinishedTimeReq struct {
	Id             int `json:"id,omitempty" valid:"id"`
	UnfinishedTime int `json:"unfinishedTime" valid:"unfinishedTime"`
}

type DeleteHabitReq struct {
	Id int `json:"id" valid:"id"`
}

func GetAllHabits(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userId").(int)
	user := &model.User{Id: userId}
	var habits []model.Habit
	var err error
	if habits, err = user.GetAllHabits(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "获取失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "获取成功",
		"habits": habits,
	})
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
	updateHabitReq := &UpdateHabitReq{}
	if err := c.ShouldBindJSON(&updateHabitReq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}
	habit := &model.Habit{
		Id:             updateHabitReq.Id,
		Name:           updateHabitReq.Name,
		Desc:           updateHabitReq.Desc,
		Difficulty:     updateHabitReq.Difficulty,
		Tag:            updateHabitReq.Tag,
		UserId:         updateHabitReq.UserId,
		FinishedTime:   updateHabitReq.FinishedTime,
		UnFinishedTime: updateHabitReq.UnFinishedTime,
	}
	if err := habit.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "更新成功",
		"habit": habit,
	})
}

func UpdateHabitFinishedTime(c *gin.Context) {
	updateUnfinishedTime := &UpdateHabitUnfinishedTimeReq{}
	if err := c.ShouldBindJSON(&updateUnfinishedTime); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}
	habit := &model.Habit{
		Id:           updateUnfinishedTime.Id,
		FinishedTime: updateUnfinishedTime.UnfinishedTime,
	}
	if err := habit.UpdateFinishedTime(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":          "更新成功",
		"finishedTime": habit.FinishedTime,
	})
}

func UpdateHabitUnfinishedTime(c *gin.Context) {
	updateHabitUnfinishedTimeReq := &UpdateHabitFinishedTimeReq{}
	if err := c.ShouldBindJSON(&updateHabitUnfinishedTimeReq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}
	habit := &model.Habit{
		Id:             updateHabitUnfinishedTimeReq.Id,
		UnFinishedTime: updateHabitUnfinishedTimeReq.FinishedTime,
	}
	if err := habit.UpdateUnfinishedTime(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":            "更新成功",
		"unfinishedTime": habit.UnFinishedTime,
	})
}

func DeleteHabit(c *gin.Context) {
	deleteHabitReq := &DeleteHabitReq{}
	if err := c.ShouldBindJSON(&deleteHabitReq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}
	habit := &model.Habit{
		Id: deleteHabitReq.Id,
	}
	if err := habit.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}
