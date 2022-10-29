package captcha

import (
	"errors"
	"go-tagle/conf"
	"go-tagle/pkg/logger"
	"go-tagle/pkg/redislib"
	"time"
)

type RedisStore struct {
	RedisClient *redislib.RedisClient
	Key         string
}

func (rs *RedisStore) Set(key string, value string) error {

	ExpireTime := time.Minute * time.Duration(conf.CaptchaConf.ExpireTime)

	if ok := rs.RedisClient.Set(rs.Key+key, value, ExpireTime); !ok {
		logger.WarnString("captcha", "redis", "无法存储图片验证码答案")
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	key = rs.Key + key
	val := rs.RedisClient.Get(key)
	if clear {
		rs.RedisClient.Del(key)
	}
	return val
}

func (rs *RedisStore) Verify(key, answer string, clear bool) bool {
	v := rs.Get(key, clear)
	return v == answer
}
