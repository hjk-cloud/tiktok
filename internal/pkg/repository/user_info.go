package repository

import (
	"errors"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/entity"
	"gorm.io/gorm"
	"sync"
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

func (*UserInfoDao) QueryUserById(id int64) (*entity.UserInfo, error) {
	var user entity.UserInfo
	gdb = NewGormDB()
	err := gdb.Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("用户id不存在")
	}
	return &user, nil
}
