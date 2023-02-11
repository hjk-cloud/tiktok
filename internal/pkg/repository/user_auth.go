package repository

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/entity"
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
	gdb.Model(&entity.UserAuth{}).Where("name = ?", name).Count(&count)
	return int(count), nil
}

func (*UserAuthDao) Register(user *entity.UserAuth) error {
	gdb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Select("id", "name", "password").Create(&user).Error; err != nil {
			return err
		}
		if err := tx.Select("id", "name").Create(entity.UserInfo{Id: user.Id, Name: user.Name}).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (*UserAuthDao) Login(name string, password string) (int64, error) {
	var user entity.UserAuth
	err := gdb.Where("name = ? AND password = ?", name, password).Take(user).Error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
