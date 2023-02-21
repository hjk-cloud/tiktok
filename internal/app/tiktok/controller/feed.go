package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	vo.Response
	VideoList []vo.VideoVO `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	var latestTime int64
	// 获取请求参数的时间
	times, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err == nil {
		// 1. 有参数，赋值
		fmt.Println("times: ", times)
		latestTime = times
	} else {
		// 2. 无参数，默认当前时间
		fmt.Println(err)
		latestTime = time.Now().Unix()
	}
	videos, nextTime, err := service.QueryFeedList(token, latestTime)
	if err == nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  vo.Response{StatusCode: 0, StatusMsg: "success"},
			VideoList: videos,
			NextTime:  nextTime,
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response: vo.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}

}
