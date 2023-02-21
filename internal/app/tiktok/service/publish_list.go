package service

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

// TODO
func PublishList(dto *dto.PublishListDTO) ([]vo.Video, error) {
	userId, err := util.JWTAuth(dto.Token)
	if err != nil {
		return nil, err
	}

	authorId := dto.UserId
	authorDo, err := repo.NewUserInfoDaoInstance().QueryUserById(authorId)
	author := vo.User{}
	if err == nil && authorDo != nil {
		author = vo.User{
			Id:            authorDo.Id,
			Name:          authorDo.Name,
			FollowCount:   authorDo.FollowCount,
			FollowerCount: authorDo.FollowerCount,
			// IsFollow:      false, // TODO: 暂时不提供，可能并没有影响
			IsFollow: getFollowStatus(userId, authorId),
		}
	}

	vos, err := repo.NewVideoRepoInstance().QueryVideoInProfile(userId, authorId)
	// vos := make([]vo.Video, len(dos))
	// if err != nil {
	// 	return vos, nil
	// }
	// // video -> favorite
	// // followDos, err := repo.NewFollowDaoInstance()
	for i := range vos {
		// vos[i] = vo.Video{
		// 	Id:            do.Id,
		// 	Author:        author,
		// 	PlayUrl:       do.PlayUrl,
		// 	CoverUrl:      do.CoverUrl,
		// 	FavoriteCount: do.FavoriteCount,
		// 	CommentCount:  do.CommentCount,
		// 	IsFavorite:    false, // TODO
		// 	Title:         do.Title,
		// }
		vos[i].Author = author
	}
	return vos, nil
}
