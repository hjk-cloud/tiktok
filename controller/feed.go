package controller

import (
	"fmt"
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
	var token string
	var latestTime int64
	// 获取请求参数的时间
	times, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err == nil {
		fmt.Println("times: ", times)
		latestTime = times
	} else {
		fmt.Println(err)
		latestTime = time.Now().Unix()
	}
	videos, nextTime, err := service.QueryFeedList(token, latestTime)
	if err == nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  model.Response{StatusCode: 0, StatusMsg: "success"},
			VideoList: videos,
			NextTime:  nextTime,
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}

}
