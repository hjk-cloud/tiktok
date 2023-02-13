package service

import (
	"fmt"
	"github.com/hjk-cloud/tiktok/dao"
	"github.com/hjk-cloud/tiktok/model"
	"log"
	"time"
)

type FeedListFlow struct {
	LatestTime time.Time
	Videos     []*model.Video
	NextTime   int64
}

// QueryFeedList Service层接口，将视频列表返回controller
func QueryFeedList(token string, latestTime time.Time) ([]model.VideoFlow, int64, error) {
	return NewFeedListFlow(token, latestTime).Do()
}

func NewFeedListFlow(token string, latestTime time.Time) *FeedListFlow {
	return &FeedListFlow{LatestTime: latestTime}
}

/**
1. 获取一组video
2. model的video需要映射到common的video，所有的video里userid要映射到user类，并加载到video类里
3. 获取nextTime
*/

func (f *FeedListFlow) Do() ([]model.VideoFlow, int64, error) {
	// 1.
	videos, err := dao.NewVideoDaoInstance().MQueryVideoByLastTime(f.LatestTime)
	if err != nil {
		log.Print(err)
	}
	// 2.
	var videoFlows []model.VideoFlow
	for i := 0; i < len(videos); i++ {
		var videoFlow model.VideoFlow
		video := videos[i]
		videoFlow = model.VideoFlow{
			Id: video.Id,
			// 空对象，待实现：需要用到user.service
			Author:        model.DemoUser,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			// 默认值，需要识别用户+是否点赞service
			IsFavorite: true,
		}
		videoFlows = append(videoFlows, videoFlow)
	}

	//3.
	//如果没有视频了
	var nextTime time.Time
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].CreateTime
	} else {
		nextTime = time.Now()
	}

	fmt.Println("feed_list:66 | Next time:", nextTime, nextTime.Unix())
	return videoFlows, nextTime.Unix(), err
}
