package service

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

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
			IsFollow:      getFollowStatus(userId, authorId),
		}
	}

	vos, err := repo.NewVideoRepoInstance().QueryVideoInProfile(userId, authorId)
	for i := range vos {
		vos[i].Author = author
	}
	return vos, nil
}
