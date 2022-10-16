package model

import (
	"fmt"
	"go-tagle/pkg/database"
)

type Habit struct {
	Id             int    `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	Name           string `json:"name" gorm:"column:name;type:varchar(255)"`
	Desc           string `json:"desc" gorm:"column:desc;type:varchar(255)"`
	Difficulty     int    `json:"difficulty" gorm:"column:difficulty;type:int"`
	Tag            string `json:"tag" gorm:"column:tag;type:varchar(255)"`
	UserId         int    `json:"user_id" gorm:"column:user_id;type:int"`
	FinishedTime   int    `json:"finished_time" gorm:"column:finished_time;type:int"`
	UnFinishedTime int    `json:"unfinished_time" gorm:"column:unfinished_time;type:int"`
}

func (h *Habit) TableName() string {
	return "habits"
}

func (h *Habit) Create() error {
	if err := database.DB.Create(h).Error; err != nil {
		fmt.Println("创建习惯失败")
		return err
	}
	return nil
}
