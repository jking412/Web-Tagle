package conf

import "go-tagle/pkg/config"

var CaptchaConf = struct {
	Height     int
	Width      int
	Length     int
	MaxSkew    float64
	DotCount   int
	ExpireTime int //分钟
	Key        string
}{}

func initCaptchaConf() {
	CaptchaConf.Height = config.LoadInt("captcha.height", 80)
	CaptchaConf.Width = config.LoadInt("captcha.width", 240)
	CaptchaConf.Length = config.LoadInt("captcha.length", 6)
	CaptchaConf.MaxSkew = config.LoadFloat64("captcha.maxSkew", 0.7)
	CaptchaConf.DotCount = config.LoadInt("captcha.dotCount", 80)
	CaptchaConf.ExpireTime = config.LoadInt("captcha.expireTime", 15)
	CaptchaConf.Key = config.LoadString("captcha.key", "")
}
