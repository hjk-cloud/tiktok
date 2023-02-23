package repository

import (
	"sync"

	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"gorm.io/gorm/clause"
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

func (*FavoriteDao) GetCountByObjectId(objectId int64, objectType string) (int64, error) {
	var count int64
	if err := Db.Model(&favorite).Where("object_id = ? AND object_type = ?", objectId, objectType).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (*FavoriteDao) GetCountBySubjectId(subjectId int64, objectType string) (int64, error) {
	var count int64
	if err := Db.Model(&favorite).Where("subject_id = ? AND object_type = ?", subjectId, objectType).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (*FavoriteDao) GetListBySubjectId(subjectId int64, objectType string) ([]do.Favorite, error) {
	var favoriteList []do.Favorite
	if err := Db.Where("subject_id = ? AND object_type = ?", subjectId, objectType).Find(&favoriteList).Error; err != nil {
		return nil, err
	}
	return favoriteList, nil
}

func (*FavoriteDao) Insert(favorite *do.Favorite) error {
	// Db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "subject_id"}, {Name: "object_id"}, {Name: "object_type"}},
	// 	DoUpdates: clause.Assignments(map[string]interface{}{"is_deleted": 0}),
	// }).Create(&favorite)
	// if err := Db.Select("subject_id", "object_id", "object_type").Create(favorite).Error; err != nil {
	// 	return err
	// }
	return Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "subject_id"}, {Name: "object_id"}, {Name: "object_type"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"is_deleted": 0, "update_time": favorite.UpdateTime}),
	}).Create(&favorite).Error
}

func (*FavoriteDao) Delete(favorite *do.Favorite) error {
	if err := Db.Model(favorite).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}
