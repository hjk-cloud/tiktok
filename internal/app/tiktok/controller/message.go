package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

// 保存到redis比较好
var tempChat = map[string][]vo.Message{}

var messageIdSequence = int64(0)

type ChatResponse struct {
	vo.Response
	MessageList []vo.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	if user, exist := UsersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(user.Id, int64(userIdB))

		atomic.AddInt64(&messageIdSequence, 1)
		curMessage := vo.Message{
			Id:         messageIdSequence,
			FromUserId: user.Id,
			ToUserId:   int64(userIdB),
			Content:    content,
			CreateTime: time.Now().Local().Unix(),
		}
		if user.Id < int64(userIdB) {
			curMessage.FirstRead = true
		} else {
			curMessage.SecondRead = true
		}

		if messages, exist := tempChat[chatKey]; exist {
			tempChat[chatKey] = append(messages, curMessage)
		} else {
			tempChat[chatKey] = []vo.Message{curMessage}
		}
		c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	user, exist := UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	userIdB, _ := strconv.Atoi(toUserId)
	chatKey := genChatKey(user.Id, int64(userIdB))

	fmt.Printf("##### from %d to %d\n", user.Id, userIdB)
	messageList := []vo.Message{}
	leftChat := []vo.Message{}
	if user.Id < int64(userIdB) {
		for _, msg := range tempChat[chatKey] {
			if msg.FirstRead {
				if !msg.SecondRead {
					leftChat = append(leftChat, msg)
				}
				continue
			}
			msg.FirstRead = true
			if !msg.SecondRead {
				leftChat = append(leftChat, msg)
			}
			messageList = append(messageList, msg)
		}
	} else {
		for _, msg := range tempChat[chatKey] {
			if msg.SecondRead {
				if !msg.FirstRead {
					leftChat = append(leftChat, msg)
				}
				continue
			}
			msg.SecondRead = true
			if !msg.FirstRead {
				leftChat = append(leftChat, msg)
			}
			messageList = append(messageList, msg)
		}
	}
	tempChat[chatKey] = leftChat
	fmt.Printf("#####%#v\n", tempChat[chatKey])
	c.JSON(http.StatusOK, ChatResponse{Response: vo.Response{StatusCode: 0}, MessageList: messageList})
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
