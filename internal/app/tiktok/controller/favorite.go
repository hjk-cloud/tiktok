package controller

import (
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

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

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userIdString := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	r := &dto.FavoriteListDTO{Token: token, UserId: userId}
	if videoList, err := service.GetFavoriteList(r); err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: vo.Response{
				StatusCode: 0,
			},
			VideoList: videoList,
		})

	}
}
