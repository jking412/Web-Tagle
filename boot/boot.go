package boot

import (
	"fmt"
	"go-tagle/conf"
	"go-tagle/model/user"
	"go-tagle/pkg/captcha"
	"go-tagle/pkg/config"
	"go-tagle/pkg/database"
	"go-tagle/pkg/logger"
	"go-tagle/pkg/redislib"
	"go-tagle/pkg/session"
)

func Initialize() {
	config.InitConfig()
	conf.InitConf()

	logger.InitLogger(conf.LogConf.Filename,
		conf.LogConf.MaxSize,
		conf.LogConf.MaxBackup,
		conf.LogConf.MaxAge,
		conf.LogConf.IsCompress,
		conf.LogConf.LogType,
		conf.LogConf.Level)

	redislib.ConnectRedis(
		fmt.Sprintf("%s:%s", conf.RedisConf.Host, conf.RedisConf.Port),
		conf.RedisConf.Username,
		conf.RedisConf.Password,
		conf.RedisConf.DB)

	database.ConnectDB(conf.DBConf.DBName)

	database.DB.AutoMigrate(&user.User{},
		&user.EmailStatus{})

	session.Init(conf.SessionConf.Secret)

	captcha.InitCaptcha()
}
