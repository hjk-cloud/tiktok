package service

import (
	"time"

	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
)

// 发送消息
func MessageAction(dto *dto.MessageActionDTO) (int64, error) {
	msg := &do.MessageDO{
		UserId:     dto.UserId,
		ToUserId:   dto.ToUserId,
		Content:    dto.MsgContent,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if _, err := repo.NewMessageRepoInstance().Create(msg); err != nil {
		return -1, err
	}
	return msg.Id, nil
}
