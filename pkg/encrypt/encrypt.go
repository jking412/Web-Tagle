package encrypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func EncryptPassword(password string) string {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println("加密失败")
		panic(err)
	}
	return string(encryptPassword)
}

func CheckPassword(password string, encryptPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	if err != nil {
		fmt.Println("密码错误")
		return false
	}
	return true
}

func IsEncrypt(password string) bool {
	return len(password) == 60
}

func GenerateSalt() string {
	rand.Seed(time.Now().Unix())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)

}
