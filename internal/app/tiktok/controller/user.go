package controller

import (
	"net/http"
	"sync/atomic"
	"strconv"

	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/gin-gonic/gin"
)

type UserInfoResponse struct {
	vo.Response
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

var UsersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}
