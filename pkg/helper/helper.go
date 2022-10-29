package helper

import (
	"math/rand"
	"strings"
	"time"
)

func ParseActiveWeekdays(activeWeekdays string) []string {
	return strings.Split(activeWeekdays, ",")
}

func ParseForm(data string) map[string]string {
	result := make(map[string]string)
	for _, v := range strings.Split(data, "&") {
		kv := strings.Split(v, "=")
		result[kv[0]] = kv[1]
	}
	return result
}

func GenerateVerifyCode() string {
	rand.Seed(time.Now().Unix())
	a := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	b := make([]rune, 6)
	for i := range b {
		b[i] = a[rand.Intn(len(a))]
	}
	return string(b)
}
