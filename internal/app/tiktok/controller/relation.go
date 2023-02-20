package controller

import (
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
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

// RelationAction no practical effect, just check if token is valid
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

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		UserList: []vo.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		UserList: []vo.User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	token := c.Query("token")
	userId, err := util.JWTAuth(token)
	if err != nil {
		writeError(c, err)
		return
	}
	var userList []vo.User
	if userId == DemoUser.Id {
		userList = []vo.User{DemoUser, ToDemoUser}
	} else {
		userList = []vo.User{ToDemoUser, DemoUser}
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		UserList: userList,
	})
}
