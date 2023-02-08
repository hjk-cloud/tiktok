package entity

import (
	"errors"
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
	"sync"
)

type UserAuth struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

func (UserAuth) TableName() string {
	return "t_user_auth"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserByName(name string) (int, error) {
	var count int64
	model.db.Model(&UserAuth{}).Where("name = ?", name).Count(&count)
	return int(count), nil
}

func (*UserDao) Register(user *UserAuth) error {
	err := model.db.Select("id", "name", "password").Create(&user).Error
	if err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}

func (*UserDao) Login(name string, password string) (int64, error) {
	var user UserAuth
	err := model.db.Where("name = ? AND password = ?", name, password).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (*UserDao) QueryUserById(id int64) (*UserAuth, error) {
	var user UserAuth
	err := model.db.Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("未查询用户id")
	}
	return &user, nil
}
