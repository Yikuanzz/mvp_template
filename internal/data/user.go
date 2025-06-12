package data

import (
	"mvp/internal/handler"
	"mvp/utils/common"
	"mvp/utils/log"
)

type User struct {
	common.BaseModel
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

// UserRepo 用户仓库
type UserRepo struct {
	data *Data
	log  log.Logger
}

// GetUser 获取用户
func (u *UserRepo) GetUser(user *handler.User) (*handler.User, error) {
	if err := u.data.db.Where("username = ?", user.Username).First(user).Error; err != nil {
		return nil, err
	}
	u.log.Info("获取用户成功", "user", user)
	return user, nil
}

// CreateUser 创建用户
func (u *UserRepo) CreateUser(user *handler.User) (*handler.User, error) {
	if err := u.data.db.Create(user).Error; err != nil {
		return nil, err
	}
	u.log.Info("创建用户成功", "user", user)
	return user, nil
}
