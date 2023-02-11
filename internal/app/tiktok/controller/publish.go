package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/param"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	// 是否登录
	var userId int64
	//if user, exist := usersLoginInfo[token]; !exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	//} else {
	//	userId = user.Id
	//}

	title := c.PostForm("title")

	data, err := c.FormFile("data")
	if err != nil {
		writeError(c, err)
		return
	}

	p := &param.PublishActionParam{Token: token, Title: title, Data: data, UserId: userId}
	if _, err := service.PublishAction(c, p); err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	// fmt.Println("###PublishList")
	// fmt.Printf("gin.Context: %+v\n", c)
	// fmt.Printf("gin.Context: %#v\n", c)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}

func writeError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, Response{
		StatusCode: 1,
		StatusMsg:  err.Error(),
	})
}
