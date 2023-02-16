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

// TODO: hook，事务
// 收取对方发给自己的未读消息
func (*MessageRepo) MessageUnreadChat(msg *(do.MessageDO)) ([]do.MessageDO, error) {
	msgs := []do.MessageDO{}
	var err error
	err = Db.Where(&do.MessageDO{UserId: msg.ToUserId, ToUserId: msg.UserId}).Where("is_read = ?", false).Find(&msgs).Error
	if err != nil {
		return msgs, err
	}
	// map 不会像struct那样自动映射？
	// Db.Where(map[string]interface{}{"UserId": msg.ToUserId, "ToUserId": msg.UserId, "IsRead": false}).Find(&ret)
	// 更新为已读
	// err = Db.Model(&msgs).Select("IsRead").Updates(&do.MessageDO{IsRead: true}).Error
	msgIds := make([]int64, len(msgs))
	for i, msg := range msgs {
		msgIds[i] = msg.Id
	}
	err = Db.Table(do.MessageDO{}.TableName()).Where("id IN (?)", msgIds).Updates(&do.MessageDO{IsRead: true}).Error
	return msgs, err
}

// 收取双方的全部消息
func (*MessageRepo) MessageChatAll(msg *(do.MessageDO)) ([]do.MessageDO, error) {
	msgs := []do.MessageDO{}
	var err error
	err = Db.Where(&do.MessageDO{UserId: msg.UserId, ToUserId: msg.ToUserId}).
		Or(&do.MessageDO{UserId: msg.ToUserId, ToUserId: msg.UserId}).Find(&msgs).Error
	if err != nil {
		return msgs, err
	}
	msgIds := make([]int64, len(msgs))
	for i, msg := range msgs {
		msgIds[i] = msg.Id
	}
	err = Db.Table(do.MessageDO{}.TableName()).Where("id IN (?)", msgIds).Updates(&do.MessageDO{IsRead: true}).Error
	return msgs, err
}
