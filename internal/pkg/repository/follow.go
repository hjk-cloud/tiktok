package repository

import (
	"log"
	"sync"

	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
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

func (*FollowDao) QueryFollowStatus(follow do.Follow) bool {
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

func (*FollowDao) GetFollowerList(objectId int64) ([]do.Follow, error) {
	var followerList []do.Follow
	if err := Db.Where("object_id = ? AND is_deleted = 0", objectId).Find(&followerList).Error; err != nil {
		return nil, err
	}
	return followerList, nil
}

func (*FollowDao) GetFollowList(subjectId int64) ([]do.Follow, error) {
	var followList []do.Follow
	if err := Db.Where("subject_id = ? AND is_deleted = 0", subjectId).Find(&followList).Error; err != nil {
		return nil, err
	}
	return followList, nil
}

func (*FollowDao) Insert(follow *do.Follow) error {
	if err := Db.Select("subject_id", "object_id").Create(follow).Error; err != nil {
		return err
	}
	return nil
}

func (*FollowDao) Delete(follow *do.Follow) error {
	if err := Db.Model(follow).Where("subject_id = ? AND object_id = ?", follow.SubjectId, follow.ObjectId).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}

// 好友列表
func (*FollowDao) GetFriendList(userId int64) ([]vo.User, error) {
	var users []vo.User
	rows, err := Db.Raw(`
SELECT 
	c.id,
	c.name,
	c.follow_count,
	c.follower_count,
	0,
	IF(c.avatar IS NULL, "", c.avatar),
	IF(c.background IS NULL, "", c.background),
	IF(c.signature IS NULL, "", c.signature),
	c.total_favorited,
	c.favorite_count,
	c.publish_count
FROM t_follow a
INNER JOIN t_follow b ON a.object_id=b.subject_id
INNER JOIN t_user_info c ON b.subject_id=c.Id AND a.is_deleted=0 AND b.is_deleted=0
WHERE a.subject_id=? AND b.object_id=?;`, userId, userId).Rows()
	if err != nil {
		return nil, err
	}
	//.Scan(&videos).Error
	for rows.Next() {
		var user vo.User
		err = rows.Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow,
			&user.Avatar, &user.BackgroundImage, &user.Signature, &user.TotalFavorited, &user.FavoriteCount, &user.WorkCount)
		if err == nil {
			users = append(users, user)
		} else {
			log.Printf("##### %#v\n", err)
		}
	}
	return users, nil
}
