package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	"net/http"
)

type UserLoginResponse struct {
	vo.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user, err := service.UserRegister(username, password)

	if err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: vo.Response{StatusCode: 0},
			UserId:   user.UserId,
			Token:    user.Token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: vo.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user, err := service.UserLogin(username, password)

	if err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: vo.Response{StatusCode: 0},
			UserId:   user.UserId,
			Token:    user.Token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: vo.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}
