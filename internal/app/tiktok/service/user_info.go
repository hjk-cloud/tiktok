package service

import (
	"errors"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
)

func GetUserInfo(r *dto.UserLoginDTO) (*do.UserInfo, error) {
	if r.UserId == 0 {
		return nil, errors.New("id为空")
	}

	userDao := repo.NewUserInfoDaoInstance()

	user, err := userDao.QueryUserById(r.UserId)
	if err != nil {
		return nil, err
	}
	//relationDao := models.NewRelationDaoInstance()
	//favoriteDao := models.NewFavoriteDaoInstance()
	//videoDao := models.NewVideoDaoInstance()
	//var totalFavorited = 0

	////关注数
	//followCount, err := relationDao.QueryRelationCountByUserId(f.UserId)
	//if err != nil {
	//	return err
	//}
	//f.FollowCount = followCount
	////粉丝数
	//followerCount, err := relationDao.QueryRelationCountByToUserId(f.UserId)
	//if err != nil {
	//	return err
	//}
	//f.FollowerCount = followerCount
	////获赞数
	//videoIds := videoDao.QueryPublishVideoList(f.UserId)
	//for i := range videoIds {
	//	totalFavorited += favoriteDao.QueryVideoFavoriteCount(videoIds[i])
	//}
	//f.TotalFavorited = totalFavorited
	////喜欢数
	//f.FavoriteCount = favoriteDao.QueryUserFavoriteCount(f.UserId)

	return user, nil
}
