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
		IsFollow:        getFollowStatus(userId, r.UserId),
		Avatar:          defaultAvatar(userInfo.Avatar),
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
		IsFollow:        getFollowStatus(subjectId, objectId),
		Avatar:          defaultAvatar(userInfo.Avatar),
		BackgroundImage: userInfo.Background,
		Signature:       userInfo.Signature,
		TotalFavorited:  userInfo.TotalFavorited,
		FavoriteCount:   userInfo.FavoriteCount,
		WorkCount:       userInfo.PublishCount,
	}
	return user, nil
}

func defaultAvatar(avatar string) string {
	return "http://www.hummingg.com/images/01cat.png"
	// if avatar == "" {
	// 	return "http://www.hummingg.com/images/01cat.png"
	// }
	// return avatar
}
