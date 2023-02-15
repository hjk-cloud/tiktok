package service

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"time"
)

// 发送消息
func MessageAction(dto *dto.MessageActionDTO) (int64, error) {
	msg := &do.MessageDO{
		UserId:     dto.UserId,
		ToUserId:   dto.ToUserId,
		Content:    dto.MsgContent,
		CreateTime: time.Now().Local(),
		UpdateTime: time.Now().Local(),
	}
	if _, err := repo.NewMessageRepoInstance().Create(msg); err != nil {
		return -1, err
	}
	return msg.Id, nil
}
