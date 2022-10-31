package core

import (
	"github.com/gin-gonic/gin"
	"go-tagle/app/requests"
	"go-tagle/model/user"
	"go-tagle/pkg/email"
	"go-tagle/pkg/session"
	"net/http"
)

func Login(c *gin.Context) (gin.H, int) {
	userLoginReq := &requests.UserLoginReq{
		Account:  c.PostForm("account"),
		Password: c.PostForm("password"),
	}
	errs := requests.ValidateUserLoginReq(userLoginReq)
	if len(errs) > 0 {
		return gin.H{
			"msg": errs,
		}, http.StatusUnprocessableEntity
	}
	_email := struct {
		email string `valid:"email"`
	}{
		email: userLoginReq.Account,
	}
	var _user *user.User
	if errs = requests.ValidateEmail(&_email); len(errs) == 0 {
		if !user.IsExistEmail(userLoginReq.Account) {
			return gin.H{
				"msg": "邮箱不存在",
			}, http.StatusBadRequest
		}
		if !user.IsActivateEmail(userLoginReq.Account) {
			return gin.H{"msg": "邮箱未激活"}, http.StatusBadRequest
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
		return gin.H{
			"msg": "登录成功",
		}, http.StatusOK
	} else {
		return gin.H{
			"msg": "账号或密码错误",
		}, http.StatusBadRequest
	}

}

func Register(c *gin.Context) (gin.H, int) {
	userRegisterReq := &requests.UserRegisterReq{
		Username:        c.PostForm("username"),
		Password:        c.PostForm("password"),
		ConfirmPassword: c.PostForm("confirm_password"),
		Email:           c.PostForm("email"),
	}
	if userRegisterReq.Password != userRegisterReq.ConfirmPassword {
		return gin.H{"msg": "两次密码不一致"}, http.StatusUnprocessableEntity
	}
	errs := requests.ValidateUserRegisterReq(userRegisterReq)
	if len(errs) > 0 {
		return gin.H{"msg": errs}, http.StatusUnprocessableEntity
	}
	_user := &user.User{
		Username: userRegisterReq.Username,
		Password: userRegisterReq.Password,
		Email:    userRegisterReq.Email,
	}
	if errs = requests.ValidateEmail(userRegisterReq.Username); len(errs) == 0 {
		return gin.H{
			"msg": "用户名不能为邮箱格式",
		}, http.StatusBadRequest
	}
	if user.IsExistUsername(_user.Username) {
		return gin.H{"msg": "用户名已存在"}, http.StatusBadRequest
	}
	if user.IsExistEmail(_user.Email) {
		return gin.H{"msg": "邮箱已存在"}, http.StatusBadRequest
	}
	if err := _user.Create(); err != nil {
		return gin.H{"msg": "注册失败"}, http.StatusBadRequest
	}
	if err := user.CreateEmailStatus(_user.Email); err != nil {
		return gin.H{"msg": "注册失败"}, http.StatusBadRequest
	}
	email.SendActivateMsg(_user.Email)
	return gin.H{
		"msg": "注册成功",
	}, http.StatusOK
}
