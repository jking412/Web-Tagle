package model

import (
	"go-tagle/pkg/database"
	"go-tagle/pkg/encrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       int     `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	Username string  `json:"username" gorm:"column:username;type:varchar(255)"`
	Password string  `json:"password" gorm:"column:password;type:varchar(255)"`
	Email    string  `json:"email" gorm:"column:email;type:varchar(255)"`
	Salt     string  `json:"-" gorm:"column:salt;type:varchar(255)"`
	Habits   []Habit `json:"habits" gorm:"foreignKey:UserId"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	if err := database.DB.Create(u).Error; err != nil {
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
//		fmt.Println("创建用户失败")
//		return err
//	}
//	return tx.Commit().Error
//}

func (u *User) GetUserById() (*User, error) {
	var user *User
	err := database.DB.Where("id = ?", u.Id).First(&user).Error
	return user, err
}

func (u *User) GetUserByUsername() (*User, error) {
	var user *User
	err := database.DB.Where("username = ?", u.Username).First(&user).Error
	return user, err
}

func (u *User) GetUserByEmail() (*User, error) {
	var user *User
	err := database.DB.Where("email = ?", u.Email).First(&user).Error
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
	return err
}

func (u *User) GetAllHabits() ([]Habit, error) {
	var habits []Habit
	err := database.DB.Model(u).Association("Habits").Find(&habits)
	return habits, err
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
