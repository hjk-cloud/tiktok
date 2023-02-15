package controller

import (
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

// UsersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var UsersLoginInfo = map[string]vo.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	vo.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	vo.Response
	User vo.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: vo.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := vo.User{
			Id:   userIdSequence,
			Name: username,
		}
		UsersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: vo.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: vo.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: vo.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
