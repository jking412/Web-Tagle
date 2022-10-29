package conf

import "go-tagle/pkg/config"

var SessionConf = struct {
	Secret     string
	ExpireTime int
}{}

func initSessionConf() {
	SessionConf.Secret = config.LoadString("session.secret", "tagle")
	SessionConf.ExpireTime = config.LoadInt("session.expireTime", 60*60*24*7)
}
