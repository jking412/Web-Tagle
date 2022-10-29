package conf

import "go-tagle/pkg/config"

var RedisConf = struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       int
}{}

var RedisDefaultExpireTime int // 分钟

func initRedisConf() {
	RedisConf.Host = config.LoadString("redis.host", "127.0.0.1")
	RedisConf.Port = config.LoadString("redis.port", "6379")
	RedisConf.Username = config.LoadString("redis.username", "")
	RedisConf.Password = config.LoadString("redis.password", "")
	RedisConf.DB = config.LoadInt("redis.db", 0)
	RedisDefaultExpireTime = config.LoadInt("redis.defaultExpireTime", 15)
}
