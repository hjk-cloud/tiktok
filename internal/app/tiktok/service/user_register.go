package service

import (
	"errors"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

type UserRegisterFlow struct {
	Username string
	Password string
	User     *do.UserAuth
	UserId   int64
	Token    string
}

func UserRegister(username string, password string) (*UserRegisterFlow, error) {
	return NewUserRegisterFlow(username, password).Do()
}

func NewUserRegisterFlow(username string, password string) *UserRegisterFlow {
	return &UserRegisterFlow{Username: username, Password: password}
}

func (f *UserRegisterFlow) Do() (*UserRegisterFlow, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.Register(); err != nil {
		return nil, err
	}
	//if err := f.packData(); err != nil {
	//	return nil, err
	//}
	return f, nil
}

func (f *UserRegisterFlow) checkParam() error {
	if f.Username == "" {
		return errors.New("用户名不能为空")
	}
	//TODO 判断非法字符输入

	userDao := repo.NewUserAuthDaoInstance()
	if count, err := userDao.QueryUserByName(f.Username); err == nil && count > 0 {
		return errors.New("用户名已存在")
	}

	return nil
}

func (f *UserRegisterFlow) Register() error {
	userDao := repo.NewUserAuthDaoInstance()

	worker := util.NewWorker(f.UserId)
	id := worker.GetId()

	password := util.Argon2Encrypt(f.Password)

	user := &do.UserAuth{
		Id:       id,
		Name:     f.Username,
		Password: password,
	}
	err := userDao.Register(user)
	if err != nil {
		return err
	}
	token, err := util.GenToken(user.Id)
	if err != nil {
		return err
	}
	f.Token = token
	f.UserId = id
	return nil
}

func (f *UserRegisterFlow) packData() error {
	f.User = &do.UserAuth{
		Id: f.UserId,
	}
	return nil
}
