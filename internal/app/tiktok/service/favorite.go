package service

import (
	"errors"
	"fmt"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

func GetFavoriteStatus(subjectId, objectId int64) bool {
	favorite := &do.Favorite{SubjectId: subjectId, ObjectId: objectId, ObjectType: "video"}
	return repo.NewFavoriteDaoInstance().QueryFavoriteStatus(favorite)
}

func UpdateFavoriteStatus(r *dto.FavoriteActionDTO) error {
	favorite := &do.Favorite{ObjectId: r.VideoId, ObjectType: "video"}
	if userId, err := util.JWTAuth(r.Token); err != nil {
		return errors.New("User doesn't exist")
	} else {
		favorite.SubjectId = userId
	}
	favoriteDao := repo.NewFavoriteDaoInstance()
	if r.ActionType {
		favoriteDao.Insert(favorite)
	} else {
		favoriteDao.Delete(favorite)
	}
	return nil
}

func GetFavoriteList(r *dto.FavoriteListDTO) ([]vo.Video, error) {

	userId, err := util.JWTAuth(r.Token)
	if err != nil {
		return nil, errors.New("User doesn't exist")
	}
	favoriteList, err := repo.NewFavoriteDaoInstance().GetListBySubjectId(userId, "video")
	if err != nil {
		return nil, err
	}
	//TODO queryVideoVOListByBatchId
	fmt.Println(favoriteList)
	var videoVOList []vo.Video

	return videoVOList, nil
}

func getFavoritedCount(objectId int64) (int64, error) {
	return repo.NewFavoriteDaoInstance().GetCountByObjectId(objectId, "video")
}
