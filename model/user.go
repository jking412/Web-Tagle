package model

import (
	"go-tagle/pkg/database"
	"go-tagle/pkg/encrypt"
	"go-tagle/pkg/logger"
	"gorm.io/gorm"
)

type User struct {
	Id       int     `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	Username string  `json:"username" gorm:"column:username;type:varchar(255)"`
	Password string  `json:"password" gorm:"column:password;type:varchar(255)"`
	Email    string  `json:"email" gorm:"column:email;type:varchar(255)"`
	Salt     string  `json:"-" gorm:"column:salt;type:varchar(255)"`
	Habits   []Habit `json:"habits" gorm:"foreignKey:UserId"`
	Tasks    []Task  `json:"tasks" gorm:"foreignKey:UserId"`
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

//func (u *User) Create() error {
//	tx := database.DB.Begin()
//	defer func() {
//		if r := recover(); r != nil {
//			tx.Rollback()
//		}
//	}()
//	if err := tx.Create(u).Error; err != nil {
//		tx.Rollback()
//		logger.WarnString("database", "创建用户失败", err.Error())
//		return err
//	}
//	return tx.Commit().Error
//}

func (u *User) GetUserById() (*User, error) {
	var user *User
	err := database.DB.Where("id = ?", u.Id).First(&user).Error
	if err != nil {
<<<<<<< HEAD
		logger.WarnString("database", "获取用户失败", err.Error())
=======
		logger.ErrorString("database", "获取用户失败", err.Error())
>>>>>>> origin/main
	}
	return user, err
}

func (u *User) GetUserByUsername() (*User, error) {
	var user *User
	err := database.DB.Where("username = ?", u.Username).First(&user).Error
<<<<<<< HEAD
	if err != nil {
		logger.WarnString("database", "获取用户失败", err.Error())
	}
=======
	logger.ErrorString("database", "获取用户失败", err.Error())
>>>>>>> origin/main
	return user, err
}

func (u *User) GetUserByEmail() (*User, error) {
	var user *User
	err := database.DB.Where("email = ?", u.Email).First(&user).Error
	if err != nil {
<<<<<<< HEAD
		logger.WarnString("database", "获取用户失败", err.Error())
=======
		logger.ErrorString("database", "获取用户失败", err.Error())
>>>>>>> origin/main
	}
	return user, err
}

func (u *User) IsExistUsername() bool {
	var count int64
	database.DB.Model(&User{}).Where("username = ?", u.Username).Count(&count)
	return count > 0
}

func (u *User) IsExistEmail() bool {
	var count int64
	database.DB.Model(&User{}).Where("email = ?", u.Email).Count(&count)
	return count > 0
}

func (u *User) DeleteUserByUsername() error {
	err := database.DB.Where("username=?", u.Username).Delete(&User{}).Error
	if err != nil {
<<<<<<< HEAD
		logger.WarnString("database", "删除用户失败", err.Error())
=======
		logger.ErrorString("database", "删除用户失败", err.Error())
>>>>>>> origin/main
	}
	return err
}

func (u *User) GetAllHabits() ([]Habit, error) {
	var habits []Habit
	err := database.DB.Model(u).Association("Habits").Find(&habits)
	if err != nil {
		logger.WarnString("database", "获取用户习惯失败", err.Error())
	}
	return habits, err
}

func (u *User) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := database.DB.Model(u).Association("Tasks").Find(&tasks)
	if err != nil {
		logger.WarnString("database", "获取用户任务失败", err.Error())
	}
	return tasks, err
}

func (u *User) CheckPassword() bool {
	var user User
	if u.Username == "" {
		database.DB.Model(&User{}).Where("email = ?", u.Email).Scan(&user)
	} else {
		database.DB.Model(&User{}).Where("username = ?", u.Username).Scan(&user)
	}
	return encrypt.CheckPassword(u.Password+user.Salt, user.Password)
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Salt == "" {
		u.Salt = encrypt.GenerateSalt()
	}
	if !encrypt.IsEncrypt(u.Password) {
		u.Password = encrypt.EncryptPassword(u.Password + u.Salt)
	}
	return
}
