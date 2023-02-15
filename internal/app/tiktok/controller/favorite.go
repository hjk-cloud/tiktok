package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
