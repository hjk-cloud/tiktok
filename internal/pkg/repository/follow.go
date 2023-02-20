package repository

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"sync"
)

type FollowDao struct {
}

var follow do.Follow
var followDao *FollowDao
var followOnce sync.Once

func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(
		func() {
			followDao = &FollowDao{}
		})
	return followDao
}

func (*FollowDao) QueryFollowStatus(follow *do.Follow) bool {
	result := Db.Where("subject_id = ? AND object_id = ?", follow.SubjectId, follow.ObjectId).Take(&follow)
	return result.RowsAffected == 1
}

func (*FollowDao) GetCountByObjectId(objectId int64) (int64, error) {
	var count int64
	if err := Db.Model(&follow).Where("object_id = ?", objectId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (*FollowDao) GetCountBySubjectId(subjectId int64) (int64, error) {
	var count int64
	if err := Db.Model(&follow).Where("subject_id = ?", subjectId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (*FollowDao) Insert(follow *do.Follow) error {
	if err := Db.Select("subject_id", "object_id").Create(follow).Error; err != nil {
		return err
	}
	return nil
}

func (*FollowDao) Delete(follow *do.Follow) error {
	if err := Db.Model(follow).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}