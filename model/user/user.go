package user

import (
	"go-tagle/model"
	"go-tagle/pkg/database"
	"go-tagle/pkg/logger"
	"time"
)

type User struct {
	Id        int           `gorm:"column:id;primaryKey;autoIncrement"`
	Username  string        `gorm:"column:username;type:varchar(32)"`
	Password  string        `gorm:"column:password;type:varchar(32)"`
	Email     string        `gorm:"column:email;type:varchar(32)"`
	AvatarUrl string        `gorm:"column:avatar_url;type:varchar(255)"`
	CreatedAt time.Time     `gorm:"column:created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at"`
	Habits    []model.Habit `gorm:"-"`
	Tasks     []model.Task  `gorm:"-"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	if err := database.DB.Create(u).Error; err != nil {
		logger.ErrorString("database", "创建用户失败", err.Error())
		return err
	}
	return nil
}
