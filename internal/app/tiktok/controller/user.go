package controller

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

type UserInfoResponse struct {
	vo.Response
	UserInfo vo.User `json:"user"`
}

func UserInfo(c *gin.Context) {
	userIdString := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	token := c.Query("token")

	r := &dto.UserLoginDTO{UserId: userId, Token: token}
	if user, err := service.GetUserInfo(r); err == nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: vo.Response{StatusCode: 0},
			UserInfo: user,
		})
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
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
