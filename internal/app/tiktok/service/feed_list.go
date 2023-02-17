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
		// 2.1 遍历视频列表
		video := videos[i]
		// 2.2 获取视频作者信息
		var userInfoDao *repository.UserInfoDao
		user, err := userInfoDao.QueryUserById(video.AuthorId)
		if err != nil {
			log.Print(err)
		}
		var userVO vo.UserVO
		userVO = vo.UserVO{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			// [TO DO] 需要关注接口
			IsFollow: false,
		}
		// 2.3 映射视频VO信息
		var videoFlow vo.VideoVO
		videoFlow = vo.VideoVO{
			Id:            video.Id,
			Author:        userVO,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			// [TO DO]: 默认值，需要识别用户+是否点赞service
			IsFavorite: false,
		}
		// 2.4
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
