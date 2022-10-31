package core

import (
	"github.com/gin-gonic/gin"
	"go-tagle/app/requests"
	"go-tagle/conf"
	"go-tagle/model/user"
	"go-tagle/pkg/captcha"
	"go-tagle/pkg/email"
	"go-tagle/pkg/redislib"
	"go-tagle/pkg/session"
	"net/http"
	"time"
)

func GenerateCaptcha(c *gin.Context) (gin.H, int) {
	id, b64s, err := captcha.GenerateCaptcha()
	if err != nil {
		return gin.H{
			"msg": "验证码生成失败",
		}, http.StatusInternalServerError
	}
	b64s = captcha.RemoveB64sPrefix(b64s)
	redislib.GlobalRedis.Set(id+id, b64s, time.Minute*time.Duration(conf.RedisDefaultExpireTime))
	return gin.H{
		"verifyCodeId": id,
		"captchaImg":   b64s,
	}, http.StatusOK
}

func SendEmailCaptcha(c *gin.Context) (gin.H, int) {
	sendEmailReq := &requests.SendEmailReq{
		Email:        c.Query("email"),
		VerifyCodeId: c.Query("verify_code_id"),
		VerifyCode:   c.Query("verify_code"),
	}
	if !user.IsExistEmail(sendEmailReq.Email) {
		return gin.H{
			"msg": "邮箱不存在",
		}, http.StatusBadRequest
	}
	if !user.IsActivateEmail(sendEmailReq.Email) {
		return gin.H{
			"msg": "邮箱未激活",
		}, http.StatusBadRequest
	}
	if !captcha.VerifyCaptcha(sendEmailReq.VerifyCodeId, sendEmailReq.VerifyCode) {
		return gin.H{
			"msg": "验证码错误",
		}, http.StatusBadRequest
	}
	email.SendVerifyCode(sendEmailReq.Email)
	b64s := redislib.GlobalRedis.Get(sendEmailReq.VerifyCodeId + sendEmailReq.VerifyCodeId)
	redislib.GlobalRedis.Del(sendEmailReq.VerifyCodeId + sendEmailReq.VerifyCodeId)
	return gin.H{
		"verifyCodeId": sendEmailReq.VerifyCodeId,
		"captchaImg":   b64s,
		"email":        sendEmailReq.Email,
	}, http.StatusOK
}

func EmailLogin(c *gin.Context) (gin.H, int) {
	emailLoginReq := &requests.EmailLoginReq{
		Email:           c.PostForm("email"),
		EmailVerifyCode: c.PostForm("email_verify_code"),
	}
	verifyCode := redislib.GlobalRedis.Get(emailLoginReq.Email)
	if verifyCode == "" || verifyCode != emailLoginReq.EmailVerifyCode {
		return gin.H{
			"msg": "验证码错误",
		}, http.StatusBadRequest
	}
	_user, ok := user.GetUserByEmail(emailLoginReq.Email)
	if !ok {
		return gin.H{
			"msg": "用户不存在",
		}, http.StatusInternalServerError
	}
	redislib.GlobalRedis.Del(emailLoginReq.Email)
	session.Save(_user.Id, c)
	return gin.H{}, http.StatusOK
}
