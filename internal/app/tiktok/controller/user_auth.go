package controller

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

type UserLoginResponse struct {
	vo.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	r := &dto.UserAuthDTO{Username: username, Password: password}
	if user, err := service.UserRegister(r); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: vo.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: vo.Response{StatusCode: 0},
			UserId:   user.UserId,
			Token:    user.Token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	r := &dto.UserAuthDTO{Username: username, Password: password}
	user, err := service.UserLogin(r)

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
