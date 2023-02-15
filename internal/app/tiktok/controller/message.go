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

var tempChat = map[string][]vo.Message{}

var messageIdSequence = int64(1)

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
			Content:    content,
			CreateTime: time.Now().Format(time.Kitchen),
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

	if user, exist := UsersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(user.Id, int64(userIdB))

		fmt.Printf("##### from %d to %d\n", user.Id, userIdB)
		tempChat[chatKey] = []vo.Message{
			vo.Message{
				Id:         1,
				Content:    "111",
				CreateTime: time.Now().Format("2006-01-02 15:04:05"),
			},
			vo.Message{
				Id:         1,
				Content:    "222",
				CreateTime: time.Now().Format("2006-01-02 15:04:05"),
			},
		}
		fmt.Printf("##### %#v\n", tempChat)

		c.JSON(http.StatusOK, ChatResponse{Response: vo.Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	} else {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
