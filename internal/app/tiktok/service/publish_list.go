package service

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
)

// TODO
func PublishList(dto *dto.PublishListDTO) ([]vo.Video, error) {
	authorDo, _ := repo.NewUserInfoDaoInstance().QueryUserById(dto.UserId)
	author := vo.User{
		Id:            authorDo.Id,
		Name:          authorDo.Name,
		FollowCount:   authorDo.FollowCount,
		FollowerCount: authorDo.FollowerCount,
		IsFollow:      false,
	}

	dos, err := repo.NewVideoRepoInstance().QueryVideoByAuthorId(dto.UserId)
	vos := make([]vo.Video, len(dos))
	if err != nil {
		return vos, nil
	}
	for i, do := range dos {
		vos[i] = vo.Video{
			Id:            do.Id,
			Author:        author,
			PlayUrl:       do.PlayUrl,
			CoverUrl:      do.CoverUrl,
			FavoriteCount: do.FavoriteCount,
			CommentCount:  do.CommentCount,
			IsFavorite:    false,
			Title:         do.Title,
		}
	}
	return vos, nil
}
