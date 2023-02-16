package service

import (
	"fmt"

	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

var (
	// ua、ub刚进入双方聊天界面时都要拉取全部聊天记录
	chats map[int64]int64
)

func init() {
	chats = make(map[int64]int64)
}

// 获取聊条记录
// 无法知道何时退出聊天窗口
func MessageChat(dto *dto.MessageChatDTO) ([]vo.Message, error) {
	fromUserId, err := util.JWTAuth(dto.Token)
	if err != nil {
		return nil, err
	}
	toUserId := dto.ToUserId
	fmt.Printf("##### from %d to %d\n", fromUserId, toUserId)
	var dtos []do.MessageDO
	if lastToUserId, ok := chats[fromUserId]; ok && lastToUserId == toUserId {
		// 已在聊天窗口，拉取最新消息
		dtos, err = repo.NewMessageRepoInstance().MessageUnreadChat(
			&do.MessageDO{UserId: fromUserId, ToUserId: toUserId})
	} else {
		// 刚进聊天窗口，拉取全部消息
		chats[fromUserId] = toUserId
		dtos, err = repo.NewMessageRepoInstance().MessageChatAll(
			&do.MessageDO{UserId: fromUserId, ToUserId: toUserId})
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
			CreateTime: dtos[i].CreateTime.Unix(),
		}
	}
	return ret, nil
}

// func genChatKey(userIdA int64, userIdB int64) string {
// 	if userIdA > userIdB {
// 		return fmt.Sprintf("%d_%d", userIdB, userIdA)
// 	}
// 	return fmt.Sprintf("%d_%d", userIdA, userIdB)
// }
