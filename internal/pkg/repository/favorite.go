package repository

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"sync"
)

type FavoriteDao struct {
}

var favorite do.Favorite
var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewFavoriteDaoInstance() *FavoriteDao {
	favoriteOnce.Do(
		func() {
			favoriteDao = &FavoriteDao{}
		})
	return favoriteDao
}

func (*FavoriteDao) QueryFavoriteStatus(favorite *do.Favorite) bool {
	result := Db.Where("subject_id = ? AND object_id = ? AND object_type = ?", favorite.SubjectId, favorite.ObjectId, favorite.ObjectType).Take(&favorite)
	return result.RowsAffected == 1
}

func (*FavoriteDao) GetCountByObjectId(objectId int64, objectType string) int64 {
	var count int64
	Db.Model(&favorite).Where("object_id = ? AND object_type = ?", objectId, objectType).Count(&count)
	return count
}

func (*FavoriteDao) GetCountBySubjectId(subjectId int64, objectType string) int64 {
	var count int64
	Db.Model(&favorite).Where("subject_id = ? AND object_type = ?", subjectId, objectType).Count(&count)
	return count
}

func (*FavoriteDao) Insert(favorite *do.Favorite) error {
	if err := Db.Select("subject_id", "object_id", "object_type").Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

func (*FavoriteDao) Delete(favorite *do.Favorite) error {
	if err := Db.Model(favorite).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}
