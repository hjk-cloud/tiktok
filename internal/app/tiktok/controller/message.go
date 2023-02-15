package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

type ChatResponse struct {
	vo.Response
	MessageList []vo.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")
	user, exist := UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	userIdB, _ := strconv.Atoi(toUserId)
	msg := &dto.MessageActionDTO{UserId: user.Id, ToUserId: int64(userIdB), MsgContent: content}
	if msgId, err := service.MessageAction(msg); err != nil {
		writeError(c, err)
		return
	} else {
		fmt.Printf("新建消息Id: %d\n", msgId)
	}

	c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
}

// MessageChat all users have same follow list
// 如此轮询：理解为聊天室功能
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	user, exist := UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	userIdB, _ := strconv.Atoi(toUserId)
	fmt.Printf("##### from %d to %d\n", user.Id, userIdB)
	msg := &dto.MessageChatDTO{UserId: user.Id, ToUserId: int64(userIdB)}
	messageList := service.MessageChat(msg)

	c.JSON(http.StatusOK, ChatResponse{Response: vo.Response{StatusCode: 0}, MessageList: messageList})
}

// func genChatKey(userIdA int64, userIdB int64) string {
// 	if userIdA > userIdB {
// 		return fmt.Sprintf("%d_%d", userIdB, userIdA)
// 	}
// 	return fmt.Sprintf("%d_%d", userIdA, userIdB)
// }
