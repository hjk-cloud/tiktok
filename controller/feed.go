package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/model"
	"github.com/hjk-cloud/tiktok/service"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//c.JSON(http.StatusOK, FeedResponse{
	//	Response:  Response{StatusCode: 0},
	//	VideoList: DemoVideos,
	//	NextTime:  time.Now().Unix(),
	//})
	latestTime := time.Now()
	videos, _ := service.QueryFeedList(latestTime)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
