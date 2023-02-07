package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/util"
)

const (
	MaxUsernameLength = 20
	MaxPasswordLength = 20
	MinPasswordLength = 6
)

type LoginFlow struct {
	Username string
	Password string
	User     *model.UserAuth
	UserId   int64
	Token    string
}

func UserLogin(username string, password string) (*LoginFlow, error) {
	return NewLoginFlow(username, password).Do()
}

func NewLoginFlow(username string, password string) *LoginFlow {
	return &LoginFlow{Username: username, Password: password}
}

func (f *LoginFlow) Do() (*LoginFlow, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareData(); err != nil {
		return nil, err
	}
	if err := f.packData(); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *LoginFlow) checkParam() error {
	if f.Username == "" {
		return errors.New("用户名为空")
	}
	if len(f.Username) > MaxUsernameLength {
		return errors.New("用户名长度超出限制")
	}
	if f.Password == "" {
		return errors.New("密码为空")
	}
	if len(f.Password) > MaxPasswordLength || len(f.Password) < MinPasswordLength {
		return errors.New("密码长度应为6-20个字符")
	}
	return nil
}

func (f *LoginFlow) prepareData() error {
	userDao := model.NewUserDaoInstance()
	password := util.Argon2Encrypt(f.Password)
	userId, err := userDao.Login(f.Username, password)
	if err != nil {
		return err
	}
	f.UserId = userId
	token, err := util.GenToken(userId)
	if err != nil {
		return err
	}
	f.Token = token
	return nil
}

func (f *LoginFlow) packData() error {
	f.User = &model.UserAuth{
		Id:   f.UserId,
		Name: f.Username,
	}
	return nil
}
