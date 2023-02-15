package service

import (
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
)

// 获取未读聊条记录
func MessageChat(dto *dto.MessageChatDTO) []vo.Message {
	dtos := repo.NewMessageRepoInstance().MessageChat(&do.MessageDO{UserId: dto.UserId, ToUserId: dto.ToUserId})
	n := len(dtos)
	ret := make([]vo.Message, n)
	for i := 0; i < n; i++ {
		ret[i] = vo.Message{
			Id:         dtos[i].Id,
			FromUserId: dtos[i].UserId,
			ToUserId:   dtos[i].ToUserId,
			Content:    dtos[i].Content,
			CreateTime: dtos[i].CreateTime.Unix(),
		}
	}
	return ret
}
