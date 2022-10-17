package model

import (
	"go-tagle/pkg/database"
	"go-tagle/pkg/logger"
)

type Habit struct {
	Id             int    `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	Name           string `json:"name" gorm:"column:name;type:varchar(255)"`
	Desc           string `json:"desc" gorm:"column:desc;type:varchar(255)"`
	Difficulty     int    `json:"difficulty" gorm:"column:difficulty;type:int"`
	Tag            string `json:"tag" gorm:"column:tag;type:varchar(255)"`
	UserId         int    `json:"userId" gorm:"column:user_id;type:int"`
	FinishedTime   int    `json:"finishedTime" gorm:"column:finished_time;type:int"`
	UnFinishedTime int    `json:"unfinishedTime" gorm:"column:unfinished_time;type:int"`
}

func (h *Habit) TableName() string {
	return "habits"
}

func (h *Habit) Create() error {
	if err := database.DB.Create(h).Error; err != nil {
		logger.ErrorString("database", "创建习惯失败", err.Error())
		return err
	}
	return nil
}

func (h *Habit) Update() error {
	if err := database.DB.Model(&Habit{}).Updates(h).Error; err != nil {
		logger.ErrorString("database", "更新习惯失败", err.Error())
		return err
	}
	return nil
}

func (h *Habit) UpdateFinishedTime() error {
	if err := database.DB.Model(&Habit{}).Where("id = ?", h.Id).Update("finished_time", h.FinishedTime).Error; err != nil {
		logger.ErrorString("database", "更新习惯失败", err.Error())
		return err
	}
	return nil
}

func (h *Habit) UpdateUnfinishedTime() error {
	if err := database.DB.Model(&Habit{}).Where("id = ?", h.Id).Update("unfinished_time", h.UnFinishedTime).Error; err != nil {
		logger.ErrorString("database", "更新习惯失败", err.Error())
		return err
	}
	return nil
}

func (h *Habit) Delete() error {
	if err := database.DB.Delete(h).Error; err != nil {
		logger.ErrorString("database", "删除习惯失败", err.Error())
		return err
	}
	return nil
}
