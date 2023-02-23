package service

import (
	"errors"
	"log"
	"time"

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
	log.Printf("##### FavoriteActionDTO %#v\n", r)
	favorite := &do.Favorite{ObjectId: r.VideoId, ObjectType: "video", CreateTime: time.Now(), UpdateTime: time.Now()}
	if userId, err := util.JWTAuth(r.Token); err != nil {
		return errors.New("User doesn't exist")
	} else {
		favorite.SubjectId = userId
	}
	video, _ := repo.NewVideoRepoInstance().QueryVideoById(r.VideoId)
	favoriteDao := repo.NewFavoriteDaoInstance()
	userDao := repo.NewUserInfoDaoInstance()
	videoDao := repo.NewVideoRepoInstance()
	if r.ActionType {
		favoriteDao.Insert(favorite)
		userDao.Add(favorite.SubjectId, "favorite_count")
		userDao.Add(video.AuthorId, "total_favorited")
		videoDao.Add(video.Id, "favorite_count")
	} else {
		favoriteDao.Delete(favorite)
		userDao.Remove(favorite.SubjectId, "favorite_count")
		userDao.Remove(video.AuthorId, "total_favorited")
		videoDao.Remove(video.Id, "favorite_count")
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
