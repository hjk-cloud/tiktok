package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

type VideoListResponse struct {
	vo.Response
	VideoList []vo.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		writeError(c, err)
		return
	}

	p := &dto.PublishActionDTO{Context: c, Token: token, Title: title, Data: data}
	if _, err := service.PublishAction(p); err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		writeError(c, err)
		return
	}
	p := &dto.PublishListDTO{Token: token, UserId: userId}
	userVideos, err := service.PublishList(p)
	if err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		VideoList: userVideos,
	})
}
