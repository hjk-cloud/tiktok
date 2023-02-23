package repository

import (
	"errors"
	"sync"

	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"gorm.io/gorm"
)

type UserInfoDao struct {
}

var userInfoDao *UserInfoDao
var userInfoOnce sync.Once

func NewUserInfoDaoInstance() *UserInfoDao {
	userInfoOnce.Do(
		func() {
			userInfoDao = &UserInfoDao{}
		})
	return userInfoDao
}

func (*UserInfoDao) QueryUserById(id int64) (*do.UserInfo, error) {
	var user do.UserInfo
	err := Db.Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("用户id不存在")
	}
	return &user, nil
}

func (*UserInfoDao) Add(id int64, column string) error {
	var user do.UserInfo
	if err := Db.Where("id = ?", id).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err := Db.Model(&user).UpdateColumn(column, gorm.Expr(column+" + 1")).Error; err != nil {
		return err
	}
	return nil
}

func (*UserInfoDao) Remove(id int64, column string) error {
	var user do.UserInfo
	if err := Db.Where("id = ?", id).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err := Db.Model(&user).UpdateColumn(column, gorm.Expr(column+" - 1")).Error; err != nil {
		return err
	}
	return nil
}
