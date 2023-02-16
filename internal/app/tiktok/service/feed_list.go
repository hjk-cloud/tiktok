package service

import (
	"fmt"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	"github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"log"
	"time"
)

type FeedService struct {
	Token      string
	LatestTime int64
	Videos     []*do.VideoDO
	NextTime   int64
}

// QueryFeedList Service层接口，将视频列表返回controller
func QueryFeedList(token string, latestTime int64) ([]vo.VideoVO, int64, error) {
	return NewFeedListFlow(token, latestTime).Do()
}

func NewFeedListFlow(token string, latestTime int64) *FeedService {
	return &FeedService{Token: token, LatestTime: latestTime}
}

/**
0. 校验参数
1. 获取一组video
2. model的video需要映射到common的video，所有的video里userid要映射到user类，并加载到video类里
3. 获取nextTime
*/

func (f *FeedService) Do() ([]vo.VideoVO, int64, error) {
	// 0. 如果时间戳超过当前时间，则等于当前时间
	if f.LatestTime > time.Now().Unix() {
		f.LatestTime = time.Now().Unix()
	}
	//fmt.Println("@@@@@Read Time: ", f.LatestTime)

	// 1.
	lt := time.Unix(f.LatestTime, 0)
	videos, err := repository.NewVideoRepoInstance().MQueryVideoByLastTime(lt)
	if err != nil {
		log.Print(err)
	}
	// 2.
	var videoFlows []vo.VideoVO
	for i := 0; i < len(videos); i++ {
		var videoFlow vo.VideoVO
		video := videos[i]
		videoFlow = vo.VideoVO{
			Id: video.Id,
			// 空对象，待实现：需要用到user.service
			Author:        repository.DemoUser,
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

	fmt.Println("feed_list:75 | Next time:", nextTime, nextTime.Unix())
	return videoFlows, nextTime.Unix(), err
}
