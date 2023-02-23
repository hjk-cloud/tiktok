package service

import (
	"errors"

	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
	"gorm.io/gorm"
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
		// TODO
		repo.Db.Model(&do.VideoDO{Id: r.VideoId}).Update("favorite_count", gorm.Expr("favorite_count + 1"))
	} else {
		favoriteDao.Delete(favorite)
		repo.Db.Model(&do.VideoDO{Id: r.VideoId}).Update("favorite_count", gorm.Expr("favorite_count - 1"))
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

	videoVOList := make([]vo.Video, len(favoriteList))
	for i := range favoriteList {
		videoVOList[i], err = getFavoriteVideo(userId, favoriteList[i].ObjectId)
	}

	return videoVOList, nil
}

func getFavoriteVideo(userId int64, videoId int64) (vo.Video, error) {
	video, err := repo.NewVideoRepoInstance().QueryVideoById(videoId)
	if err != nil {
		return vo.Video{}, err
	}
	videoVO := vo.Video{
		Id:            video.Id,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		IsFavorite:    GetFavoriteStatus(userId, video.Id),
	}
	return videoVO, nil
}
