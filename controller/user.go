package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-tagle/model"
	"go-tagle/pkg/validate"
	"net/http"
)

type UserRegisterReq struct {
	Username string `valid:"username" json:"username"`
	Password string `valid:"password" json:"password"`
	Email    string `valid:"email" json:"email"`
}

type UserLoginReq struct {
	Account  string `json:"account" valid:"account"`
	Password string `json:"password" valid:"password"`
}

func Register(c *gin.Context) {
	userRegisterReq := &UserRegisterReq{}
	if err := c.ShouldBindJSON(&userRegisterReq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}
	errs := validate.ValidateUserRegisterReq(userRegisterReq)
	if len(errs) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": errs,
		})
		return
	}
	user := &model.User{
		Username: userRegisterReq.Username,
		Password: userRegisterReq.Password,
		Email:    userRegisterReq.Email,
	}
	if user.IsExistUsername() {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名已存在",
		})
		return
	}
	if user.IsExistEmail() {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "邮箱已存在",
		})
		return
	}
	if err := user.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "注册失败",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("userId", user.Id)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功",
		"user": user,
	})
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userId")
	if userId != nil {
		user, err := (&model.User{
			Id: userId.(int),
		}).GetByUserId()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "已登录",
				"user": user,
			})
			return
		}
	}
	userLoginReq := &UserLoginReq{}
	if err := c.ShouldBindJSON(&userLoginReq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "JSON格式错误",
		})
		return
	}
	errs := validate.ValidateUserLoginReq(userLoginReq)
	if len(errs) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": errs,
		})
		return
	}
	email := struct {
		email string `valid:"email"`
	}{
		email: userLoginReq.Account,
	}
	var user *model.User
	if errs = validate.ValidateEmail(&email); len(errs) == 0 {
		user = &model.User{
			Email:    userLoginReq.Account,
			Password: userLoginReq.Password,
		}
	} else {
		user = &model.User{
			Username: userLoginReq.Account,
			Password: userLoginReq.Password,
		}
	}
	if ok := user.CheckPassword(); ok {
		session.Set("userId", user.Id)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"msg":  "登录成功",
			"user": user,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "账号或密码错误",
		})
		return
	}
}
