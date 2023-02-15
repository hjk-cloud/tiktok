package controller

import (
	"errors"
	"net/http"

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
	// TODO：检查是否登录，放入service会导致循环引用，UsersLoginInfo不应该放controller？
	var userId int64
	if user, exist := UsersLoginInfo[token]; !exist {
		writeError(c, errors.New("User doesn't exist"))
		return
	} else {
		userId = user.Id
	}

	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		writeError(c, err)
		return
	}

	p := &dto.PublishActionDTO{Context: c, Token: token, Title: title, Data: data, UserId: userId}
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
	c.JSON(http.StatusOK, VideoListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
