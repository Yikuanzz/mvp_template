package handler

import (
	"fmt"

	"mvp/utils/log"

	"mvp/utils/response"

	"github.com/gin-gonic/gin"
)

// User 用户结构体
type User struct {
	ID       string
	Username string
	Password string
}

// UserRepo 用户仓库接口
type UserRepo interface {
	GetUser(user *User) (*User, error)
	CreateUser(user *User) (*User, error)
}

// UserHandler 用户处理程序
type UserHandler struct {
	userRepo UserRepo
	log      log.Logger
}

// Login 用户登录
func (u *UserHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, fmt.Errorf("参数错误: %v", err).Error())
		return
	}

	user := &User{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := u.userRepo.GetUser(user)
	if err != nil {
		response.Error(c, 400, fmt.Errorf("用户不存在: %v", err).Error())
		return
	}

	u.log.Info("[Login] 用户 %v 登录成功", user.Username)

	response.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}

// Register 用户注册
func (u *UserHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, fmt.Errorf("参数错误: %v", err).Error())
		return
	}

	user := &User{
		Username: req.Username,
		Password: req.Password,
	}

	_, err := u.userRepo.CreateUser(user)
	if err != nil {
		response.Error(c, 400, fmt.Errorf("注册失败: %v", err).Error())
		return
	}

	response.Success(c, nil)
}
