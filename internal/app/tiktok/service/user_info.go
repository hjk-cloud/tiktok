package service

import (
	"errors"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

func GetUserInfo(r *dto.UserLoginDTO) (vo.User, error) {
	var user vo.User
	if r.UserId == 0 {
		return user, errors.New("id为空")
	}

	userId, err := util.JWTAuth(r.Token)

	userDao := repo.NewUserInfoDaoInstance()

	userInfo, err := userDao.QueryUserById(r.UserId)
	if err != nil {
		return user, err
	}
	user = vo.User{
		Id:              userInfo.Id,
		Name:            userInfo.Name,
		FollowCount:     userInfo.FollowCount,
		FollowerCount:   userInfo.FollowerCount,
		IsFollow:        GetFollowStatus(userId, r.UserId),
		Avatar:          userInfo.Avatar,
		BackgroundImage: userInfo.Background,
		Signature:       userInfo.Signature,
		TotalFavorited:  userInfo.TotalFavorited,
		FavoriteCount:   userInfo.FavoriteCount,
		WorkCount:       userInfo.PublishCount,
	}
	return user, nil
}

func getUserInfoById(subjectId, objectId int64) (vo.User, error) {
	var user vo.User
	userDao := repo.NewUserInfoDaoInstance()

	userInfo, err := userDao.QueryUserById(objectId)
	if err != nil {
		return user, err
	}
	user = vo.User{
		Id:              userInfo.Id,
		Name:            userInfo.Name,
		FollowCount:     userInfo.FollowCount,
		FollowerCount:   userInfo.FollowerCount,
		IsFollow:        GetFollowStatus(subjectId, objectId),
		Avatar:          userInfo.Avatar,
		BackgroundImage: userInfo.Background,
		Signature:       userInfo.Signature,
		TotalFavorited:  userInfo.TotalFavorited,
		FavoriteCount:   userInfo.FavoriteCount,
		WorkCount:       userInfo.PublishCount,
	}
	return user, nil
}
