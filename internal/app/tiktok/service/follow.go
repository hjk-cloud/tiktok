package service

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

func getFollowStatus(subjectId, objectId int64) bool {
	if subjectId == objectId {
		return false
	}
	follow := do.Follow{SubjectId: subjectId, ObjectId: objectId}
	return repo.NewFollowDaoInstance().QueryFollowStatus(follow)
}

func UpdateFollowStatus(r *dto.FollowActionDTO) error {
	follow := &do.Follow{ObjectId: r.ToUserId}
	if userId, err := util.JWTAuth(r.Token); err != nil {
		return err
	} else {
		follow.SubjectId = userId
	}
	followDao := repo.NewFollowDaoInstance()
	if r.ActionType {
		followDao.Insert(follow)
	} else {
		followDao.Delete(follow)
	}
	return nil
}

func GetFollowList(r *dto.FollowRelationDTO) ([]vo.User, error) {
	userId, err := util.JWTAuth(r.Token)
	if err != nil {
		return nil, err
	}

	followDao := repo.NewFollowDaoInstance()
	followList, err := followDao.GetFollowList(userId)
	if err != nil {
		return nil, err
	}
	user := make([]vo.User, len(followList))
	for i := range followList {
		user[i], err = getUserInfoById(userId, followList[i].ObjectId)
	}
	return user, nil
}

func GetFollowerList(r *dto.FollowRelationDTO) ([]vo.User, error) {
	userId, err := util.JWTAuth(r.Token)
	if err != nil {
		return nil, err
	}

	followDao := repo.NewFollowDaoInstance()
	followList, err := followDao.GetFollowerList(r.UserId)
	if len(followList) == 0 {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	user := make([]vo.User, len(followList))
	for i := range followList {
		user[i], err = getUserInfoById(followList[i].SubjectId, userId)
	}
	return user, nil
}
