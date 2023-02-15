package repository

import (
	"errors"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"gorm.io/gorm"
	"sync"
)

type UserAuthDao struct {
}

var userAuthDao *UserAuthDao
var userAuthOnce sync.Once

func NewUserAuthDaoInstance() *UserAuthDao {
	userAuthOnce.Do(
		func() {
			userAuthDao = &UserAuthDao{}
		})
	return userAuthDao
}

func (*UserAuthDao) QueryUserByName(name string) (int, error) {
	var count int64
	Db.Model(&do.UserAuth{}).Where("name = ?", name).Count(&count)
	return int(count), nil
}

func (*UserAuthDao) Register(user *do.UserAuth) error {
	Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Select("id", "name", "password").Create(&user).Error; err != nil {
			return err
		}
		if err := tx.Select("id", "name").Create(do.UserInfo{Id: user.Id, Name: user.Name}).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (*UserAuthDao) Login(name string, password string) (int64, error) {
	var user do.UserAuth
	err := Db.Where("name = ? AND password = ?", name, password).Take(user).Error
	if err == gorm.ErrRecordNotFound {
		return 0, errors.New("用户名或密码错误")
	}
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
