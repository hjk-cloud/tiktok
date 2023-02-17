package repository

import (
	"errors"
	"github.com/hjk-cloud/tiktok/config"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"

	//_ "github.com/spf13/viper" // 只执行init()
	"gorm.io/gorm"
)

// TableName why？
// type Video entity.Video

// type Video struct {
// 	entity.Video
// }

// func NewGormDB() *gorm.DB {
// 	// sync.once
// 	getConfig()
// 	log.Println("#####", connLink)
// 	db, err := gorm.Open(mysql.Open(connLink))
// 	if err != nil {
// 		log.Panicln(err)
// 		// panic(err)
// 	}
// 	return db
// }

type VideoRepo struct {
}

var videoRepo *VideoRepo
var videoOnce sync.Once

func NewVideoRepoInstance() *VideoRepo {
	videoOnce.Do(
		func() {
			videoRepo = &VideoRepo{}
		})
	return videoRepo
}

// 奇怪:太短
func (*VideoRepo) Create(video *(do.VideoDO)) (int64, error) {
	// fmt.Println("#####", video.TableName())
	// gdb = NewGormDB()
	err := Db.Create(video).Error
	return video.Id, err
}

func (*VideoRepo) ExistUidHash(userid int64, hash string) bool {
	// fmt.Println("#####", userid, hash)
	// gdb = NewGormDB()
	var video *do.VideoDO
	result := Db.Where(&do.VideoDO{AuthorId: userid, HashValue: hash}, // , Status: 0
		map[string]interface{}{"Status": 0}).Find(video)
	// fmt.Println(result)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// QueryVideoById 通过id查询单个video
func (*VideoRepo) QueryVideoById(id int) (*do.VideoDO, error) {
	var video do.VideoDO
	// 主要查询语句：查询单个元素语句，错误当场处理
	err := Db.Where("id = ?", id).First(&video).Error

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

// MQueryVideoByLastTime 根据时间倒序查询多个video
func (*VideoRepo) MQueryVideoByLastTime(latestTime time.Time) ([]do.VideoDO, error) {
	var videos []do.VideoDO
	// 查询语句：查询符合条件的元素，错误当场处理
	err := Db.Where("create_time < ?", latestTime).
		Order("create_time desc").
		Limit(config.FeedMaxNum).
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
