package encrypt

import (
	"go-tagle/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logger.WarnString("encrypt", "加密密码失败", err.Error())
		panic(err)
	}
	return string(encryptPassword)
}

func CheckPassword(password string, encryptPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	if err != nil {
		logger.InfoString("encrypt", "密码错误", err.Error())
		return false
	}
	return true
}

func IsEncrypt(password string) bool {
	return len(password) == 60
}
