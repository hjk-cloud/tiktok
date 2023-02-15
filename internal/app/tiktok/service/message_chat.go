package service

import (
	"fmt"

	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

// 获取未读聊条记录
func MessageChat(dto *dto.MessageChatDTO) ([]vo.Message, error) {
	userId, err := util.JWTAuth(dto.Token)
	if err != nil {
		return nil, err
	}
	fmt.Printf("##### from %d to %d\n", userId, dto.ToUserId)
	dtos, err := repo.NewMessageRepoInstance().MessageChat(&do.MessageDO{UserId: userId, ToUserId: dto.ToUserId})
	if err != nil {
		return nil, err
	}
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
	return ret, nil
}
