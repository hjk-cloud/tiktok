package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/app/tiktok/service"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	"github.com/hjk-cloud/tiktok/util"
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
	userId, err := util.JWTAuth(token)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	userIdB, _ := strconv.Atoi(toUserId)
	msg := &dto.MessageActionDTO{UserId: userId, ToUserId: int64(userIdB), MsgContent: content}
	if msgId, err := service.MessageAction(msg); err != nil {
		writeError(c, err)
		return
	} else {
		fmt.Printf("新建消息Id: %d\n", msgId)
	}

	c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
}

// MessageChat all users have same follow list
// 如此轮询，理解为聊天室功能
// 且无法知道何时退出了聊天室
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	userIdB, err := strconv.Atoi(toUserId)
	if err != nil {
		writeError(c, err)
		return
	}
	msg := &dto.MessageChatDTO{Token: token, ToUserId: int64(userIdB)}
	messageList, err := service.MessageChat(msg)
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, ChatResponse{Response: vo.Response{StatusCode: 0}, MessageList: messageList})
}
