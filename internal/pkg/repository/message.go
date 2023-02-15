package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"sync"
)

type MessageRepo struct {
}

var msgRepo *MessageRepo
var msgOnce sync.Once

func NewMessageRepoInstance() *MessageRepo {
	msgOnce.Do(
		func() {
			msgRepo = &MessageRepo{}
		})
	return msgRepo
}

func (*MessageRepo) Create(msg *(do.MessageDO)) (int64, error) {
	err := Db.Create(msg).Error
	return msg.Id, err
}

// 收取发给自己的未读消息
func (*MessageRepo) MessageChat(msg *(do.MessageDO)) []do.MessageDO {
	ret := []do.MessageDO{}
	Db.Where(&do.MessageDO{UserId: msg.ToUserId, ToUserId: msg.UserId}).Where("is_read = ?", false).Find(&ret)
	// map 不会像struct那样自动映射？
	// Db.Where(map[string]interface{}{"UserId": msg.ToUserId, "ToUserId": msg.UserId, "IsRead": false}).Find(&ret)
	return ret
}
