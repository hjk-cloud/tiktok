package repository

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/hjk-cloud/tiktok/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"

	"gorm.io/gorm"
)

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

func (*VideoRepo) Create(video *(do.VideoDO)) (int64, error) {
	err := Db.Create(video).Error
	return video.Id, err
}

func (*VideoRepo) ExistUidHash(userId int64, hash string) bool {
	log.Println("#####", userId, hash)
	result := Db.Where(&do.VideoDO{AuthorId: userId, HashValue: hash},
		map[string]interface{}{"Status": 0}).First(&do.VideoDO{})
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// QueryVideoById 通过id查询单个video
func (*VideoRepo) QueryVideoById(id int64) (*do.VideoDO, error) {
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

// 用户访问作者(也可能是用户本人)主页
func (*VideoRepo) QueryVideoInProfile(userId int64, authorId int64) ([]vo.Video, error) {
	var videos []vo.Video
	rows, err := Db.Raw(`
	SELECT a.Id, a.play_url, a.cover_url, a.comment_count, a.favorite_count, a.title, 1-IF(b.is_deleted IS NULL, 1, b.is_deleted) AS is_follow 
	FROM t_video a LEFT JOIN t_favorite b 
	ON a.id=b.object_Id AND b.subject_id=? AND b.is_deleted=0 AND b.object_type='video'
	WHERE a.author_id=? AND a.Status=0;`, userId, authorId).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var video vo.Video
		err = rows.Scan(&video.Id, &video.PlayUrl, &video.CoverUrl, &video.CommentCount, &video.FavoriteCount,
			&video.Title, &video.IsFavorite)
		if err == nil {
			videos = append(videos, video)
		} else {
			log.Printf("##### %#v\n", err)
		}
	}
	return videos, rows.Err()
}

func (*VideoRepo) Add(id int64, column string) error {
	var video do.VideoDO
	if err := Db.Where("id = ?", id).Take(&video).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err := Db.Model(&video).UpdateColumn(column, gorm.Expr(column+" + 1")).Error; err != nil {
		return err
	}
	return nil
}

func (*VideoRepo) Remove(id int64, column string) error {
	var video do.VideoDO
	if err := Db.Where("id = ?", id).Take(&video).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err := Db.Model(&video).UpdateColumn(column, gorm.Expr(column+" - 1")).Error; err != nil {
		return err
	}
	return nil
}
