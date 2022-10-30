package core

import (
	"github.com/gin-gonic/gin"
	"go-tagle/controller/requests"
	"go-tagle/model/user"
	"go-tagle/pkg/email"
	"go-tagle/pkg/session"
	"net/http"
)

func Login(c *gin.Context) (interface{}, int) {
	userLoginReq := &requests.UserLoginReq{
		Account:  c.PostForm("account"),
		Password: c.PostForm("password"),
	}
	errs := requests.ValidateUserLoginReq(userLoginReq)
	if len(errs) > 0 {
		return errs, http.StatusUnprocessableEntity
	}
	_email := struct {
		email string `valid:"email"`
	}{
		email: userLoginReq.Account,
	}
	var _user *user.User
	if errs = requests.ValidateEmail(&_email); len(errs) == 0 {
		if !user.IsExistEmail(userLoginReq.Account) {
			return "邮箱不存在", http.StatusBadRequest
		}
		if !user.IsActivateEmail(userLoginReq.Account) {
			return "邮箱未激活", http.StatusBadRequest
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
		return nil, http.StatusOK
	} else {
		return "账号或密码错误", http.StatusBadRequest
	}

}

func Register(c *gin.Context) (interface{}, int) {
	userRegisterReq := &requests.UserRegisterReq{
		Username:        c.PostForm("username"),
		Password:        c.PostForm("password"),
		ConfirmPassword: c.PostForm("confirm_password"),
		Email:           c.PostForm("email"),
	}
	if userRegisterReq.Password != userRegisterReq.ConfirmPassword {
		return "两次密码不一致", http.StatusUnprocessableEntity
	}
	errs := requests.ValidateUserRegisterReq(userRegisterReq)
	if len(errs) > 0 {
		return errs, http.StatusUnprocessableEntity
	}
	_user := &user.User{
		Username: userRegisterReq.Username,
		Password: userRegisterReq.Password,
		Email:    userRegisterReq.Email,
	}
	if errs = requests.ValidateEmail(userRegisterReq.Username); len(errs) == 0 {
		return "用户名不能为邮箱格式", http.StatusBadRequest
	}
	if user.IsExistUsername(_user.Username) {
		return "用户名已存在", http.StatusBadRequest
	}
	if user.IsExistEmail(_user.Email) {
		return "邮箱已存在", http.StatusBadRequest
	}
	if err := _user.Create(); err != nil {
		return "注册失败", http.StatusBadRequest
	}
	if err := user.CreateEmailStatus(_user.Email); err != nil {
		return "注册失败", http.StatusBadRequest
	}
	email.SendActivateMsg(_user.Email)
	return nil, http.StatusOK
}
