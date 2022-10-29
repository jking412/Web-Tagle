package captcha

import (
	"github.com/mojocn/base64Captcha"
	"go-tagle/conf"
	"go-tagle/pkg/logger"
	"go-tagle/pkg/redislib"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

var internalCaptcha *Captcha

func InitCaptcha() {
	once.Do(func() {
		internalCaptcha = &Captcha{
			Base64Captcha: newCaptcha(),
		}
		logger.InfoString("captcha", "初始化成功", "")
	})
}

func newCaptcha() *base64Captcha.Captcha {

	store := RedisStore{
		RedisClient: redislib.GlobalRedis,
		Key:         conf.CaptchaConf.Key,
	}

	driver := base64Captcha.NewDriverDigit(conf.CaptchaConf.Height,
		conf.CaptchaConf.Width,
		conf.CaptchaConf.Length,
		conf.CaptchaConf.MaxSkew,
		conf.CaptchaConf.DotCount)

	return base64Captcha.NewCaptcha(driver, &store)
}

func RemoveB64sPrefix(b64s string) string {
	return b64s[len("data:image/png;base64,"):]
}

func GenerateCaptcha() (id string, b64s string, err error) {
	return internalCaptcha.Base64Captcha.Generate()
}

func VerifyCaptcha(id string, answer string) (match bool) {
	return internalCaptcha.Base64Captcha.Verify(id, answer, true)
}
