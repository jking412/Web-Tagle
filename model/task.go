package model

import (
	"go-tagle/pkg/database"
	"go-tagle/pkg/logger"
)

type Task struct {
	Id             int    `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	Name           string `json:"name" gorm:"column:name;type:varchar(255)"`
	Desc           string `json:"desc" gorm:"column:desc;type:varchar(255)"`
	Difficulty     int    `json:"difficulty" gorm:"column:difficulty;type:int"`
	Tag            string `json:"tag" gorm:"column:tag;type:varchar(255)"`
	ActiveWeekdays string `json:"activeWeekdays" gorm:"column:active_weekdays;type:varchar(255)"`
	UserId         int    `json:"userId" gorm:"column:user_id;type:int"`
	FinishedTime   int    `json:"finishedTime" gorm:"column:finished_time;type:int"`
}

func (t *Task) TableName() string {
	return "tasks"
}

func (t *Task) Create() error {
	if err := database.DB.Create(t).Error; err != nil {
		logger.WarnString("database", "创建任务失败", err.Error())
		return err
	}
	return nil
}

func (t *Task) Update() error {
	if err := database.DB.Model(&Task{}).Updates(t).Error; err != nil {
		logger.WarnString("database", "更新任务失败", err.Error())
		return err
	}
	return nil
}

func (t *Task) UpdateFinishedTime() error {
	if err := database.DB.Model(&Task{}).Where("id = ?", t.Id).Update("finished_time", t.FinishedTime).Error; err != nil {
		logger.WarnString("database", "更新任务失败", err.Error())
		return err
	}
	return nil
}

func (t *Task) Delete() error {
	if err := database.DB.Delete(t).Error; err != nil {
		logger.WarnString("database", "删除任务失败", err.Error())
		return err
	}
	return nil
}
