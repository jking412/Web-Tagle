package user

import (
	"go-tagle/model"
	"go-tagle/pkg/database"
	"go-tagle/pkg/encrypt"
	"go-tagle/pkg/logger"
)

func GetAllHabits(u *User) ([]model.Habit, error) {
	var habits []model.Habit
	err := database.DB.Model(u).Association("Habits").Find(&habits)
	if err != nil {
		logger.WarnString("database", "获取用户习惯失败", err.Error())
	}
	return habits, err
}

func GetAllTasks(u *User) ([]model.Task, error) {
	var tasks []model.Task
	err := database.DB.Model(u).Association("Tasks").Find(&tasks)
	if err != nil {
		logger.WarnString("database", "获取用户任务失败", err.Error())
	}
	return tasks, err
}

func GetUserById(id int) (*User, bool) {
	var user *User
	err := database.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.WarnString("database", "获取用户失败", err.Error())
		return nil, false
	}
	return user, true
}

func GetUserByUsername(username string) (*User, error) {
	var user *User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.WarnString("database", "获取用户失败", err.Error())
	}
	return user, err
}

func GetUserByEmail(email string) (*User, bool) {
	user := &User{}
	err := database.DB.Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		logger.WarnString("database", "获取用户失败", err.Error())
		return nil, false
	}
	return user, true
}

func IsExistUsername(username string) bool {
	var count int64
	database.DB.Model(&User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func IsExistEmail(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func DeleteUserByUsername(username string) error {
	err := database.DB.Where("username=?", username).Delete(&User{}).Error
	if err != nil {
		logger.WarnString("database", "删除用户失败", err.Error())
	}
	return err
}

func CreateEmailStatus(email string) error {
	var emailStatus EmailStatus
	emailStatus.Email = email
	emailStatus.IsActivate = false
	err := database.DB.Create(&emailStatus).Error
	return err
}

func IsActivateEmail(email string) bool {
	var emailStatus EmailStatus
	database.DB.Model(&EmailStatus{}).Where("email = ?", email).Scan(&emailStatus)
	return emailStatus.IsActivate
}

func ActivateEmail(email string) bool {
	if err := database.DB.Table("email_info").Where("email = ?", email).Update("is_activate", true).Error; err != nil {
		logger.WarnString("database", "激活邮箱失败", err.Error())
		return false
	}
	return true
}

func CheckPassword(u *User) bool {
	var user User
	if u.Username == "" {
		database.DB.Model(&User{}).Where("email = ?", u.Email).Scan(&user)
	} else {
		database.DB.Model(&User{}).Where("username = ?", u.Username).Scan(&user)
	}
	return encrypt.CheckPassword(u.Password, user.Password)
}
