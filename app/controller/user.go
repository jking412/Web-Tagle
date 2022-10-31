package controller

import (
	"github.com/gin-gonic/gin"
	"go-tagle/app/core"
	"go-tagle/conf"
	"go-tagle/model/user"
	"go-tagle/pkg/oauth2"
	"go-tagle/pkg/session"
	"net/http"
)

type UserController struct {
}

func (uc *UserController) Register(c *gin.Context) {
	_, ok := session.GetUser(c)
	if ok {
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.HTML(http.StatusOK, "register", gin.H{})
}

func (uc *UserController) DoRegister(c *gin.Context) {
	data, statusCode := core.Register(c)
	if statusCode == http.StatusOK {
		c.Redirect(http.StatusFound, "/user/login")
	} else {
		c.HTML(statusCode, "error", data)
	}
}

func (uc *UserController) Login(c *gin.Context) {
	_, ok := session.GetUser(c)
	if ok {
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.HTML(http.StatusOK, "login", gin.H{})
}

func (uc *UserController) DoLogin(c *gin.Context) {
	data, statusCode := core.Login(c)
	if statusCode == http.StatusOK {
		c.Redirect(http.StatusFound, "/")
	} else {
		c.HTML(statusCode, "error", data)
	}
}

func (uc *UserController) GithubLogin(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://github.com/login/oauth/authorize?client_id="+
		conf.GithubClientConf.ClientID)
}

func (uc *UserController) GithubLoginCallback(c *gin.Context) {
	code := c.Query("code")
	githubUserInfo, ok := oauth2.GetGithubUserInfo(code)
	if !ok {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"msg": "获取github用户信息失败",
		})
		return
	}
	_user := &user.User{
		Username:  githubUserInfo.Username,
		AvatarUrl: githubUserInfo.AvatarUrl,
	}
	var err error

	if flag := user.IsExistUsername(_user.Username); !flag {
		if err = _user.Create(); err != nil {
			c.HTML(http.StatusInternalServerError, "error", gin.H{
				"msg": "注册失败",
			})
			return
		}
	}

	if _user, err = user.GetUserByUsername(_user.Username); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"msg": "获取用户信息失败",
		})
		return
	}
	session.Save(_user.Id, c)
	c.Redirect(http.StatusFound, "/")
}
