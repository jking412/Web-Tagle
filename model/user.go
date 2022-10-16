package model

import (
	"fmt"
	"go-tagle/pkg/database"
	"go-tagle/pkg/encrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `json:"id;omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Email    string `json:"email" gorm:"column:email"`
	Salt     string `json:"-" gorm:"column:salt"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	if err := database.DB.Create(u).Error; err != nil {
		fmt.Println("创建用户失败")
		return err
	}
	return nil
}

func (u *User) GetByUserId() (User, error) {
	var user User
	err := database.DB.Where("id = ?", u.Id).First(&user).Error
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
