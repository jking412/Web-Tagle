package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-tagle/conf"
	"go-tagle/model/user"
	"go-tagle/pkg/email"
	"go-tagle/pkg/helper"
	"go-tagle/pkg/oauth2"
	"go-tagle/pkg/request"
	"go-tagle/pkg/session"
	"io/ioutil"
	"net/http"
	"strings"
)

type UserController struct {
}

type UserRegisterReq struct {
	Username        string `valid:"username" form:"username"`
	Password        string `valid:"password" form:"password"`
	ConfirmPassword string `form:"confirm_password"`
	Email           string `valid:"email" form:"email"`
}

type UserLoginReq struct {
	Account  string `form:"account" valid:"account"`
	Password string `form:"password" valid:"password"`
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
	userRegisterReq := &UserRegisterReq{
		Username:        c.PostForm("username"),
		Password:        c.PostForm("password"),
		ConfirmPassword: c.PostForm("confirm_password"),
		Email:           c.PostForm("email"),
	}
	if userRegisterReq.Password != userRegisterReq.ConfirmPassword {
		c.HTML(http.StatusUnprocessableEntity, "error", gin.H{
			"msg": "两次密码不一致",
		})
	}
	errs := request.ValidateUserRegisterReq(userRegisterReq)
	if len(errs) > 0 {
		c.HTML(http.StatusUnprocessableEntity, "error", gin.H{
			"msg": errs,
		})
		return
	}
	_user := &user.User{
		Username: userRegisterReq.Username,
		Password: userRegisterReq.Password,
		Email:    userRegisterReq.Email,
	}
	if errs = request.ValidateEmail(userRegisterReq.Username); len(errs) == 0 {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "用户名不能为邮箱格式",
		})
	}
	if user.IsExistUsername(_user.Username) {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "用户名已存在",
		})
		return
	}
	if user.IsExistEmail(_user.Email) {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "邮箱已存在",
		})
		return
	}
	if err := _user.Create(); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"msg": "注册失败",
		})
		return
	}
	if err := user.CreateEmailStatus(_user.Email); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"msg": "注册失败",
		})
		return
	}
	email.SendActivateMsg(_user.Email)
	c.Redirect(http.StatusFound, "/user/login")
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
	userLoginReq := &UserLoginReq{
		Account:  c.PostForm("account"),
		Password: c.PostForm("password"),
	}
	errs := request.ValidateUserLoginReq(userLoginReq)
	if len(errs) > 0 {
		c.HTML(http.StatusUnprocessableEntity, "error", gin.H{
			"msg": errs,
		})
		return
	}
	_email := struct {
		email string `valid:"email"`
	}{
		email: userLoginReq.Account,
	}
	var _user *user.User
	if errs = request.ValidateEmail(&_email); len(errs) == 0 {
		if !user.IsExistEmail(userLoginReq.Account) {
			c.HTML(http.StatusBadRequest, "error", gin.H{
				"msg": "邮箱不存在",
			})
			return
		}
		if !user.IsActivateEmail(userLoginReq.Account) {
			c.HTML(http.StatusBadRequest, "error", gin.H{
				"msg": "邮箱未激活",
			})
			return
		}
		_user = &user.User{
			Email:    userLoginReq.Account,
			Password: userLoginReq.Password,
		}
	} else {
		_user = &user.User{
			Username: userLoginReq.Account,
			Password: userLoginReq.Password,
		}
	}
	if ok := user.CheckPassword(_user); ok {
		if _user.Username == "" {
			_user, _ = user.GetUserByEmail(_user.Email)
		} else {
			_user, _ = user.GetUserByUsername(_user.Username)
		}
		session.Save(_user.Id, c)
		c.Redirect(http.StatusFound, "/")
		return
	} else {
		c.HTML(http.StatusBadRequest, "error", gin.H{
			"msg": "账号或密码错误",
		})
		return
	}
}

func (uc *UserController) GithubLogin(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://github.com/login/oauth/authorize?client_id="+
		conf.GithubClientConf.ClientID)
}

func (uc *UserController) GithubLoginCallback(c *gin.Context) {
	code := c.Query("code")

	resp, err := http.Post("https://github.com/login/oauth/access_token",
		"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("client_id=%s&client_secret=%s&code=%s",
			conf.GithubClientConf.ClientID,
			conf.GithubClientConf.ClientSecret,
			code)))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var token oauth2.Token
	form := helper.ParseForm(string(body))
	for k, v := range form {
		if k == "access_token" {
			token.AccessToken = v
		}
	}
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, _ = http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	var githubUserInfo oauth2.GithubUserInfo
	if err = json.Unmarshal(body, &githubUserInfo); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"msg": "Github登录失败",
		})
		return
	}
	_user := &user.User{
		Username:  githubUserInfo.Username,
		AvatarUrl: githubUserInfo.AvatarUrl,
	}

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
