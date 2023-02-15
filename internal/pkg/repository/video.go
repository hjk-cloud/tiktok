package repository

import (
	"errors"
	"sync"

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
	Db.Create(video)
	return video.Id, nil
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
