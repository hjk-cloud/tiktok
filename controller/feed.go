package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/model"
	"github.com/hjk-cloud/tiktok/service"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	model.Response
	VideoList []model.VideoFlow `json:"video_list,omitempty"`
	NextTime  int64             `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//c.JSON(http.StatusOK, FeedResponse{
	//	Response:  Response{StatusCode: 0},
	//	VideoList: DemoVideos,
	//	NextTime:  time.Now().Unix(),
	//})

	//token := c.Param("token")
	var token string
	var latestTime time.Time
	// 获取请求参数的时间
	times, err := strconv.ParseInt(c.Param("latest_time"), 10, 64)
	// 有参数用参数，无参数用当前时间
	if err == nil {
		latestTime = time.Unix(0, times*1e6).Local()
	} else {
		latestTime = time.Now()
	}

	videos, nextTime, _ := service.QueryFeedList(token, latestTime)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0, StatusMsg: "success"},
		VideoList: videos,
		NextTime:  nextTime,
	})
}
