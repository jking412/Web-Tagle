package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var Store sessions.Store

func Init() {
	Store = cookie.NewStore([]byte("secret"))
}
