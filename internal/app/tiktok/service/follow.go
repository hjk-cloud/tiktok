package service

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

func GetFollowStatus(subjectId, objectId int64) bool {
	follow := &do.Follow{SubjectId: subjectId, ObjectId: objectId}
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

//func GetFollowList()

//func GetFollowerList()

func getFollowCount(subjectId int64) (int64, error) {
	if count, err := repo.NewFollowDaoInstance().GetCountBySubjectId(subjectId); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

func GetFollowerCount(objectId int64) (int64, error) {
	if count, err := repo.NewFollowDaoInstance().GetCountByObjectId(objectId); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}
