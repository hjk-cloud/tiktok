package service

import (
	"log"
	"time"

	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

func MessageChat(dto *dto.MessageChatDTO) ([]vo.Message, error) {
	fromUserId, err := util.JWTAuth(dto.Token)
	if err != nil {
		return nil, err
	}
	toUserId := dto.ToUserId
	log.Printf("##### %d %d from %d to %d\n", dto.PreMsgTime, time.Now().Unix(), fromUserId, toUserId)
	var dtos []do.MessageDO
	if dto.PreMsgTime == 0 {
		// log.Println("##### 刚进聊天窗口")
		dtos, err = repo.NewMessageRepoInstance().MessageChatAll(
			&do.MessageDO{UserId: fromUserId, ToUserId: toUserId})
	} else {
		// log.Println("##### 已在聊天窗口", util.UnixMilliToTime(dto.PreMsgTime))
		dtos, err = repo.NewMessageRepoInstance().MessageUnreadChat(
			&do.MessageDO{UserId: fromUserId, ToUserId: toUserId, CreateTime: util.UnixMilliToTime(dto.PreMsgTime)})
	}
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
			CreateTime: dtos[i].CreateTime.UnixMilli(), // 毫秒数
		}
	}
	return ret, nil
}
