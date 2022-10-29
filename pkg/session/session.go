package session

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-tagle/conf"
	"go-tagle/model/user"
	"go-tagle/pkg/logger"
	"go-tagle/pkg/redislib"
	"strconv"
	"sync"
	"time"
)

var once sync.Once

var Store sessions.Store

func Init(secret string) {
	once.Do(func() {
		Store = cookie.NewStore([]byte(secret))
		logger.InfoString("session", "session初始化成功", "")
	})

}

func Save(userId int, c *gin.Context) {
	session := sessions.Default(c)
	session.Set("userId", userId)
	session.Options(sessions.Options{
		HttpOnly: true,
		MaxAge:   conf.SessionConf.ExpireTime,
	})
	session.Save()
}

func GetUser(c *gin.Context) (*user.User, bool) {
	session := sessions.Default(c)
	userData := session.Get("userId")
	if userData == nil {
		return nil, false
	}
	userId := userData.(int)
	userStr := redislib.GlobalRedis.Get(strconv.FormatInt(int64(userId), 10))
	_user := &user.User{}
	if userStr == "" {
		var ok bool
		if _user, ok = user.GetUserById(userId); !ok {
			logger.WarnString("session", "获取用户失败", "")
			return nil, false
		}
		data, err := json.Marshal(_user)
		if err != nil {
			logger.WarnString("session", "json序列化失败", err.Error())
			return nil, false
		} else {
			redislib.GlobalRedis.Set(strconv.FormatInt(int64(userId), 10), string(data), time.Hour*time.Duration(24))
		}
	}
	if err := json.Unmarshal([]byte(userStr), _user); err != nil {
		logger.WarnString("session", "json解析用户失败", err.Error())
		return nil, false
	}
	return _user, true
}
