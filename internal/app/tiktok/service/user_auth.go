package service

import (
	"errors"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

const (
	MaxUsernameLength = 30
	MaxPasswordLength = 20
	MinPasswordLength = 6
)

func UserRegister(r *dto.UserAuthDTO) (*vo.UserAuth, error) {
	if r.Username == "" {
		return nil, errors.New("用户名不能为空")
	}
	//TODO 判断非法字符输入

	userDao := repo.NewUserAuthDaoInstance()
	if count, err := userDao.QueryUserByName(r.Username); err == nil && count > 0 {
		return nil, errors.New("用户名已存在")
	}

	id := util.NewWorker(1).GetId()
	password := util.Argon2Encrypt(r.Password)

	user := &do.UserAuth{
		Id:       id,
		Name:     r.Username,
		Password: password,
	}
	if err := userDao.Register(user); err != nil {
		return nil, err
	}
	if token, err := util.GenToken(user.Id); err != nil {
		return nil, err
	} else {
		return &vo.UserAuth{UserId: user.Id, Token: token}, nil
	}
}

func UserLogin(r *dto.UserAuthDTO) (*vo.UserAuth, error) {
	if r.Username == "" {
		return nil, errors.New("用户名为空")
	}
	if len(r.Username) > MaxUsernameLength {
		return nil, errors.New("用户名长度超出限制")
	}
	if r.Password == "" {
		return nil, errors.New("密码为空")
	}
	if len(r.Password) > MaxPasswordLength || len(r.Password) < MinPasswordLength {
		return nil, errors.New("密码长度应为6-20个字符")
	}

	userDao := repo.NewUserAuthDaoInstance()
	password := util.Argon2Encrypt(r.Password)
	userId, err := userDao.Login(r.Username, password)
	if err != nil {
		return nil, err
	}
	token, err := util.GenToken(userId)
	if err != nil {
		return nil, err
	}
	return &vo.UserAuth{UserId: userId, Token: token}, nil
}
