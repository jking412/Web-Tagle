package conf

import "go-tagle/pkg/config"

var EmailConf = struct {
	Host                 string
	Port                 int
	Username             string
	Password             string
	VerifyCodeExpireTime int // 分钟
}{}

var EmailActivateConf = struct {
	ActivateUrl string
	ExpireTime  int // 小时
}{}

func initEmailConf() {
	EmailConf.Host = config.LoadString("email.host", "smtp.qq.com")
	EmailConf.Port = config.LoadInt("email.port", 25)
	EmailConf.Username = config.LoadString("email.username", "")
	EmailConf.Password = config.LoadString("email.password", "")
	EmailConf.VerifyCodeExpireTime = config.LoadInt("email.expireTime", 15)
	EmailActivateConf.ActivateUrl = config.LoadString("email.activateUrl",
		"http://localhost:8000/user/email/activate?email=%s&&code=%s")
	EmailActivateConf.ExpireTime = config.LoadInt("email.expireTime", 24)
}
