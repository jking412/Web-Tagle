package user

import (
	"go-tagle/pkg/encrypt"
	"gorm.io/gorm"
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if !encrypt.IsEncrypt(u.Password) {
		u.Password = encrypt.EncryptPassword(u.Password)
	}
	return
}
