package service

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

func GetFollowStatus(subjectId, objectId int64) bool {
	follow := do.Follow{SubjectId: subjectId, ObjectId: objectId}
	return repo.NewFollowDaoInstance().QueryFollowStatus(follow)
}

func UpdateFollowStatus(r *dto.FollowActionDTO) error {
	follow := &do.Follow{ObjectId: r.ToUserId}
	if userId, err := util.JWTAuth(r.Token); err != nil {
		return err
	} else {
		follow.ObjectId = userId
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
	followList, err := followDao.GetFollowList(r.UserId)
	if err != nil {
		return nil, err
	}
	var user []vo.User
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
	if err != nil {
		return nil, err
	}
	var user []vo.User
	for i := range followList {
		user[i], err = getUserInfoById(followList[i].SubjectId, userId)
	}
	return user, nil
}

func getFollowCount(subjectId int64) (int64, error) {
	if count, err := repo.NewFollowDaoInstance().GetCountBySubjectId(subjectId); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

func getFollowerCount(objectId int64) (int64, error) {
	if count, err := repo.NewFollowDaoInstance().GetCountByObjectId(objectId); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}
