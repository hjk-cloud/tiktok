package repository

import (
	"errors"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/entity"

	"gorm.io/gorm"
	"sync"
)

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
	mydb.Model(&entity.UserAuth{}).Where("name = ?", name).Count(&count)
	return int(count), nil
}

func (*UserDao) Register(user *entity.UserAuth) error {
	err := mydb.Select("id", "name", "password").Create(&user).Error
	if err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}

func (*UserDao) Login(name string, password string) (int64, error) {
	var user entity.UserAuth
	err := mydb.Where("name = ? AND password = ?", name, password).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (*UserDao) QueryUserById(id int64) (*entity.UserAuth, error) {
	var user entity.UserAuth
	err := mydb.Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("未查询用户id")
	}
	return &user, nil
}

