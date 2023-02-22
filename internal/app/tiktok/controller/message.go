package controller

import (
	"errors"
	"log"
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
		log.Printf("新建消息Id: %d\n", msgId)
	}

	c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	preMsgTime := c.Query("pre_msg_time") // 刚进窗口是0，后来是上次最新消息的时间（毫秒）
	lastTime, err := strconv.ParseInt(preMsgTime, 10, 64)
	// 秒是10位数, 毫秒是13位数。客户端传过来单位不一致？
	// if lastTime < int64(math.Pow10(12)) {
	// 	lastTime *= 1000
	// }
	if err != nil {
		writeError(c, errors.New("Invalid pre_msg_time"))
		return
	}
	log.Println("#####", lastTime, "客户端上消息列表中最后一条消息的时间")
	userIdB, err := strconv.Atoi(toUserId)
	if err != nil {
		writeError(c, err)
		return
	}
	msg := &dto.MessageChatDTO{Token: token, ToUserId: int64(userIdB), PreMsgTime: lastTime}
	messageList, err := service.MessageChat(msg)
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, ChatResponse{Response: vo.Response{StatusCode: 0}, MessageList: messageList})
}
