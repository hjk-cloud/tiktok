package controller

import (
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoIdString := c.Query("video_id")
	actionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(videoIdString, 10, 64)
	r := &dto.FavoriteActionDTO{Token: token, VideoId: videoId, ActionType: actionType == "1"}

	if err := service.UpdateFavoriteStatus(r); err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: vo.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "success"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		VideoList: repository.DemoVideos,
	})
}
