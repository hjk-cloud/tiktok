package dao

import (
	"github.com/hjk-cloud/tiktok/configs"
	"github.com/hjk-cloud/tiktok/model"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

// QueryVideoById 通过id查询单个video
func (*VideoDao) QueryVideoById(id int) (*model.Video, error) {
	var video model.Video
	// 主要查询语句：查询单个元素语句，错误当场处理
	err := model.Db.Where("id = ?", id).First(&video).Error

	// 错误1：查询不到相关元素：返回空对象，不继续抛出错误
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	// 错误2：查询出错：返回空对象，抛出错误
	if err != nil {
		log.Print("QueryUserById err: ", err)
		return nil, err
	}
	return &video, nil
}

func (*VideoDao) MQueryVideoByLastTime(latestTime time.Time) ([]model.Video, error) {
	var videos []model.Video
	// 查询语句：查询符合条件的元素，错误当场处理
	err := model.Db.Where("create_time < ?", latestTime).
		Order("create_time desc").
		Limit(configs.FeedMaxNum).
		Find(&videos).Error

	// 错误1：查询不到相关元素：返回空对象，不继续抛出错误
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	// 错误2：查询出错：返回空对象，抛出错误
	if err != nil {
		log.Print("MQueryVideoByLastTime err: ", err)
		return nil, err
	}
	return videos, nil
}
