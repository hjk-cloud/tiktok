package service

import (
	"errors"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
)

func GetUserInfo(r *dto.UserLoginDTO) (vo.User, error) {
	var user vo.User
	if r.UserId == 0 {
		return user, errors.New("id为空")
	}

	userDao := repo.NewUserInfoDaoInstance()

	userInfo, err := userDao.QueryUserById(r.UserId)
	if err != nil {
		return user, err
	}
	user = vo.User{Id: userInfo.Id, Name: userInfo.Name, FavoriteCount: userInfo.FavoriteCount, FollowCount: userInfo.FollowCount, FollowerCount: userInfo.FollowerCount, WorkCount: userInfo.PublishCount}
	return user, nil
}

func GetUserInfoById(subjectId, objectId int64) (vo.User, error) {
	var user vo.User
	userDao := repo.NewUserInfoDaoInstance()

	userInfo, err := userDao.QueryUserById(objectId)
	if err != nil {
		return user, err
	}
	user = vo.User{Id: userInfo.Id, Name: userInfo.Name, FavoriteCount: userInfo.FavoriteCount, FollowCount: userInfo.FollowCount, FollowerCount: userInfo.FollowerCount, WorkCount: userInfo.PublishCount}
	user.IsFollow = GetFollowStatus(subjectId, objectId)
	return user, nil
}
