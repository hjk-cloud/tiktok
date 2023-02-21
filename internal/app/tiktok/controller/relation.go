package controller

import (
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	"github.com/hjk-cloud/tiktok/util"
)

type UserListResponse struct {
	vo.Response
	UserList []vo.User `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserIdString := c.Query("to_user_id")
	actionType := c.Query("action_type")
	toUserId, _ := strconv.ParseInt(toUserIdString, 10, 64)
	r := &dto.FollowActionDTO{Token: token, ToUserId: toUserId, ActionType: actionType == "1"}

	if err := service.UpdateFollowStatus(r); err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
	} else {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "success"})
	}
}

func FollowList(c *gin.Context) {
	userIdString := c.Query("user_id")
	token := c.Query("token")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)

	r := &dto.FollowRelationDTO{UserId: userId, Token: token}
	if userList, err := service.GetFollowList(r); err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: vo.Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	}
}

// 粉丝列表
func FollowerList(c *gin.Context) {
	userIdString := c.Query("user_id")
	token := c.Query("token")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)

	r := &dto.FollowRelationDTO{UserId: userId, Token: token}
	if userList, err := service.GetFollowerList(r); err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: vo.Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	}
}

// 好友列表：已关注
func FriendList(c *gin.Context) {
	token := c.Query("token")
	userIdStr := c.Query("user_id")
	userId, err := util.JWTAuth(token)
	log.Printf("##### %s %s %d\n", userIdStr, token, userId)
	if err != nil {
		writeError(c, err)
		return
	}
	// var userList []vo.User
	r := &dto.FollowRelationDTO{UserId: userId, Token: token}
	if userList, err := service.GetFollowList(r); err != nil {
		// c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
		writeError(c, err)
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: vo.Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	}
}
