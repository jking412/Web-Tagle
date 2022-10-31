package core

import (
	"github.com/gin-gonic/gin"
	"go-tagle/conf"
	"go-tagle/pkg/captcha"
	"go-tagle/pkg/redislib"
	"net/http"
	"time"
)

func GenerateCaptcha(c *gin.Context) (interface{}, int) {
	data := make(gin.H)
	id, b64s, err := captcha.GenerateCaptcha()
	if err != nil {
		data["msg"] = "验证码生成失败"
		return data, http.StatusInternalServerError
	}
	b64s = captcha.RemoveB64sPrefix(b64s)
	redislib.GlobalRedis.Set(id+id, b64s, time.Minute*time.Duration(conf.RedisDefaultExpireTime))
	data = gin.H{
		"verifyCodeId": id,
		"captchaImg":   b64s,
	}
	return data, http.StatusOK
}
