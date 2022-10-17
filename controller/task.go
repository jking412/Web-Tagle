package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-tagle/model"
	"net/http"
)

type TaskCreateReq struct {
	Name string `json:"name" valid:"name"`
}

type TaskUpdateReq struct {
	Name           string `json:"name" valid:"name"`
	Desc           string `json:"desc" valid:"desc"`
	Difficulty     int    `json:"difficulty" valid:"difficulty"`
	Tag            string `json:"tag" valid:"tag"`
	ActiveWeekdays string `json:"activeWeekdays" valid:"activeWeekdays"`
}

type TaskUpdateFinishedTimeReq struct {
	Id           int `json:"id" valid:"id"`
	FinishedTime int `json:"finishedTime" valid:"finishedTime"`
}

type TaskDeleteReq struct {
	Id int `json:"id" valid:"id"`
}

func GetAllTasks(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userId").(int)
	user := &model.User{Id: userId}
	tasks, err := user.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "获取失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "获取成功",
		"tasks": tasks,
	})
}

func CreateTask(c *gin.Context) {
	createTaskReq := &TaskCreateReq{}
	if err := c.ShouldBindJSON(&createTaskReq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}

	session := sessions.Default(c)
	userId := session.Get("userId").(int)
	task := &model.Task{
		Name:   createTaskReq.Name,
		UserId: userId,
	}
	if err := task.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功",
	})
}

func UpdateTask(c *gin.Context) {
	updateTaskReq := &TaskUpdateReq{}
	if err := c.ShouldBindJSON(&updateTaskReq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}

	session := sessions.Default(c)
	userId := session.Get("userId").(int)
	task := &model.Task{
		Name:           updateTaskReq.Name,
		Desc:           updateTaskReq.Desc,
		Difficulty:     updateTaskReq.Difficulty,
		Tag:            updateTaskReq.Tag,
		ActiveWeekdays: updateTaskReq.ActiveWeekdays,
		UserId:         userId,
	}
	if err := task.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

func UpdateFinishedTime(c *gin.Context) {
	updateFinishedTimeReq := &TaskUpdateFinishedTimeReq{}
	if err := c.ShouldBindJSON(&updateFinishedTimeReq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}

	task := &model.Task{
		Id:           updateFinishedTimeReq.Id,
		FinishedTime: updateFinishedTimeReq.FinishedTime,
	}

	if err := task.UpdateFinishedTime(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":          "更新成功",
		"finishedTime": task.FinishedTime,
	})

}

func DeleteTask(c *gin.Context) {
	deleteTaskReq := &TaskDeleteReq{}
	if err := c.ShouldBindJSON(&deleteTaskReq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}

	task := &model.Task{
		Id: deleteTaskReq.Id,
	}
	if err := task.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}