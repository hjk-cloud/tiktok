package service

import (
	"fmt"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	"github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
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
func QueryFeedList(token string, latestTime int64) ([]vo.Video, int64, error) {
	return NewFeedListFlow(token, latestTime).Do()
}

func NewFeedListFlow(token string, latestTime int64) *FeedService {
	return &FeedService{Token: token, LatestTime: latestTime}
}

/*
*
0. 校验参数
1. 获取一组video
2. model的video需要映射到common的video，所有的video里userid要映射到user类，并加载到video类里
3. 获取nextTime
*/
// 全写在Do()过于臃肿，后续有空将函数拆开
func (f *FeedService) validateParams() {
}

func (f *FeedService) Do() ([]vo.Video, int64, error) {
	// 0. 如果时间戳超过当前时间，则等于当前时间
	if f.LatestTime > time.Now().Unix() {
		f.LatestTime = time.Now().Unix()
	}
	// token校验，并获取本账户ID
	fmt.Println("@@@@@token: " + f.Token)
	userId, err := util.JWTAuth(f.Token)

	// 1.
	lt := time.Unix(f.LatestTime, 0)
	videos, err := repository.NewVideoRepoInstance().MQueryVideoByLastTime(lt)
	if err != nil {
		log.Print(err)
	}
	// 2.
	var videoFlows []vo.Video
	for i := 0; i < len(videos); i++ {
		// 2.1 遍历视频列表
		video := videos[i]
		//2.2 获取视频作者信息
		var userInfoDao *repository.UserInfoDao
		var authorVO vo.User
		var isFavorite bool

		author, err := userInfoDao.QueryUserById(video.AuthorId)
		if err != nil {
			log.Print(err)
		}
		// 若user不为空，则赋值
		if author != nil {
			authorVO, err = getUserInfoById(userId, author.Id)
			if err != nil {
				log.Print(err)
			}
			//authorVO = vo.User{
			//	Id:            author.Id,
			//	Name:          author.Name,
			//	FollowCount:   author.FollowCount,
			//	FollowerCount: author.FollowerCount,
			//	// [TO DO] 需要关注接口
			//	IsFollow: getFollowStatus(userId, author.Id),
			//}
			// 需要在author不为空的时候赋值，否则author.Id会报空指针异常
			isFavorite = GetFavoriteStatus(userId, author.Id)
		}

		// 2.3 映射视频VO信息
		var videoFlow vo.Video
		videoFlow = vo.Video{
			Id:     video.Id,
			Author: authorVO,
			//Author:        repository.DemoUser,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			// [TO DO]: 默认值，需要识别用户+是否点赞service
			IsFavorite: isFavorite,
			Title:      video.Title,
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
