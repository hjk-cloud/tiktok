package repository

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
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

// 收取对方发给自己的未读消息
func (*MessageRepo) MessageUnreadChat(msg *(do.MessageDO)) ([]do.MessageDO, error) {
	msgs := []do.MessageDO{}
	var err error
	err = Db.Where(&do.MessageDO{UserId: msg.ToUserId, ToUserId: msg.UserId}).Where("create_time > ?", msg.CreateTime).Find(&msgs).Error
	return msgs, err
}

// 收取双方的全部消息
func (*MessageRepo) MessageChatAll(msgDo *(do.MessageDO)) ([]do.MessageDO, error) {
	msgs := []do.MessageDO{}
	var err error
	err = Db.Where(&do.MessageDO{UserId: msgDo.UserId, ToUserId: msgDo.ToUserId}).
		Or(&do.MessageDO{UserId: msgDo.ToUserId, ToUserId: msgDo.UserId}).Find(&msgs).Error
	return msgs, err
}
