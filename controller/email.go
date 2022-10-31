package controller

import (
	"github.com/gin-gonic/gin"
	"go-tagle/controller/core"
	"go-tagle/controller/requests"
	"go-tagle/model/user"
	"go-tagle/pkg/captcha"
	"go-tagle/pkg/email"
	"go-tagle/pkg/redislib"
	"go-tagle/pkg/session"
	"net/http"
)

func (uc *UserController) EmailLogin(c *gin.Context) {
	_, ok := session.GetUser(c)
	if ok {
		c.Redirect(http.StatusFound, "/")
		return
	}
	data, statusCode := core.GenerateCaptcha(c)
	c.HTML(statusCode, "login-email", data)
}

func (uc *UserController) SendEmailVerifyCode(c *gin.Context) {
	sendEmailReq := &requests.SendEmailReq{
		Email:        c.Query("email"),
		VerifyCodeId: c.Query("verify_code_id"),
		VerifyCode:   c.Query("verify_code"),
	}
	if !user.IsExistEmail(sendEmailReq.Email) {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "邮箱不存在",
		})
		return
	}
	if !user.IsActivateEmail(sendEmailReq.Email) {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "邮箱未激活",
		})
		return
	}
	if !captcha.VerifyCaptcha(sendEmailReq.VerifyCodeId, sendEmailReq.VerifyCode) {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "验证码错误",
		})
		return
	}
	email.SendVerifyCode(sendEmailReq.Email)
	b64s := redislib.GlobalRedis.Get(sendEmailReq.VerifyCodeId + sendEmailReq.VerifyCodeId)
	redislib.GlobalRedis.Del(sendEmailReq.VerifyCodeId + sendEmailReq.VerifyCodeId)
	c.HTML(http.StatusOK, "login-email", gin.H{
		"verifyCodeId": sendEmailReq.VerifyCodeId,
		"captchaImg":   b64s,
		"email":        sendEmailReq.Email,
	})
}

func (uc *UserController) DoEmailLogin(c *gin.Context) {
	emailLoginReq := &requests.EmailLoginReq{
		Email:           c.PostForm("email"),
		EmailVerifyCode: c.PostForm("email_verify_code"),
	}
	verifyCode := redislib.GlobalRedis.Get(emailLoginReq.Email)
	if verifyCode == "" || verifyCode != emailLoginReq.EmailVerifyCode {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "验证码错误",
		})
		return
	}
	_user, ok := user.GetUserByEmail(emailLoginReq.Email)
	if !ok {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"msg": "用户不存在",
		})
		return
	}
	redislib.GlobalRedis.Del(emailLoginReq.Email)
	session.Save(_user.Id, c)
	c.Redirect(http.StatusFound, "/")
}

func (uc *UserController) ActivateEmail(c *gin.Context) {
	_email := c.Query("email")
	code := c.Query("code")
	if code == "" {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "链接无效",
		})
		return
	}
	redisCode := redislib.GlobalRedis.Get(_email)
	if redisCode == "" || redisCode != code {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "链接无效",
		})
		return
	}
	if ok := user.ActivateEmail(_email); !ok {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"msg": "激活失败",
		})
		return
	}
	redislib.GlobalRedis.Del(_email)
	c.Redirect(http.StatusFound, "/")
}
