package controller

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
)

type UserInfoResponse struct {
	Response
	UserInfo do.UserInfo `json:"user"`
}

func UserInfo(c *gin.Context) {
	userIdString := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	token := c.Query("token")

	if user, err := service.GetUserInfo(token, userId); err == nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 0},
			UserInfo: *user,
		})
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

var UsersLoginInfo = map[string]vo.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}
