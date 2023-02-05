package service

import (
	"github.com/hjk-cloud/tiktok/dao"
	"github.com/hjk-cloud/tiktok/model"
	"log"
	"time"
)

// 有啥用？
type FeedListFlow struct {
	Videos   []*model.Video
	nextTime int64
}

// QueryFeedList Service层接口，将视频列表返回controller
func QueryFeedList(lastTime time.Time) ([]model.Video, error) {
	videos, err := dao.NewVideoDaoInstance().MQueryVideoByLastTime(lastTime)
	if err != nil {
		log.Print(err)
	}
	return videos, nil
}
