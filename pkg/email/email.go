package email

import (
	"crypto/tls"
	"fmt"
	"github.com/google/uuid"
	"go-tagle/conf"
	"go-tagle/pkg/helper"
	"go-tagle/pkg/logger"
	"go-tagle/pkg/redislib"
	"gopkg.in/gomail.v2"
	"time"
)

type EmailInfo struct {
	To      string
	Subject string
	Body    string
}

func Send(emailInfo *EmailInfo) bool {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.EmailConf.Username)
	m.SetHeader("To", emailInfo.To)
	m.SetHeader("Subject", emailInfo.Subject)
	m.SetBody("text/html", emailInfo.Body)

	d := gomail.NewDialer(conf.EmailConf.Host,
		conf.EmailConf.Port,
		conf.EmailConf.Username,
		conf.EmailConf.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		logger.WarnString("email", "send email failed", err.Error())
		return false
	}
	return true
}

func SendActivateMsg(email string) bool {
	activateCode := uuid.NewString()
	if ok := redislib.GlobalRedis.Set(email, activateCode, time.Hour*time.Duration(conf.EmailActivateConf.ExpireTime)); !ok {
		logger.WarnString("email", "set activate code failed", "")
		return false
	}
	activateUrl := fmt.Sprintf(conf.EmailActivateConf.ActivateUrl,
		email,
		activateCode)
	emailInfo := &EmailInfo{
		To:      email,
		Subject: "邮箱激活",
		Body: `<p>欢迎使用tagle，激活你的邮箱，马上开始使用</p><br>
		<a href="` + activateUrl + `">点击激活</a>`,
	}

	return Send(emailInfo)
}

func SendVerifyCode(email string) bool {
	verifyCode := helper.GenerateVerifyCode()
	if ok := redislib.GlobalRedis.Set(email, verifyCode, time.Minute*time.Duration(conf.EmailConf.VerifyCodeExpireTime)); !ok {
		logger.WarnString("email", "set verify code failed", "")
		return false
	}
	emailInfo := &EmailInfo{
		To:      email,
		Subject: "邮箱验证码",
		Body:    "您的验证码是：" + verifyCode + "，有效期为" + fmt.Sprintf("%d", conf.EmailConf.VerifyCodeExpireTime) + "分钟",
	}
	return Send(emailInfo)
}
