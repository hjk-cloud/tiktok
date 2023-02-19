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

// var (
// 	// ua、ub刚进入双方聊天界面时都要拉取全部聊天记录
// 	// 再记录一下最后一次轮询时间，超时不轮询则已退出窗口
// 	// 经测试，客户端约每1300ms轮询一次
// 	// fromUserId: [toUserId, lastQueryTime]
// 	chats map[int64][2]int64
// 	// 超过这个时间没轮询就是退出聊天窗口
// 	timeLimit int64
// )

// func init() {
// 	timeLimit = int64(config.Config.ChatPollingInterval) * 1000 // 七秒钟的记忆
// 	chats = make(map[int64][2]int64)
// }

// // 获取聊条记录
// // 无法知道何时退出聊天窗口
// func MessageChatOrigin(dto *dto.MessageChatDTO) ([]vo.Message, error) {
// 	log.Println("##### MessageChat")
// 	fromUserId, err := util.JWTAuth(dto.Token)
// 	if err != nil {
// 		return nil, err
// 	}
// 	toUserId := dto.ToUserId
// 	log.Printf("##### %d from %d to %d\n", time.Now().UnixMilli(), fromUserId, toUserId)
// 	var dtos []do.MessageDO
// 	if lastToUser, ok := chats[fromUserId]; ok && lastToUser[0] == toUserId &&
// 		lastToUser[1]+timeLimit > time.Now().UnixMilli() {
// 		// 已在聊天窗口，拉取最新消息
// 		log.Println("##### 已在聊天窗口")
// 		// chats[fromUserId][1] = time.Now().UnixMilli()	// 报错
// 		chats[fromUserId] = [2]int64{toUserId, time.Now().UnixMilli()}
// 		dtos, err = repo.NewMessageRepoInstance().MessageUnreadChat(
// 			&do.MessageDO{UserId: fromUserId, ToUserId: toUserId})
// 	} else {
// 		// 刚进聊天窗口，拉取全部消息
// 		log.Println("##### 刚进聊天窗口")
// 		chats[fromUserId] = [2]int64{toUserId, time.Now().UnixMilli()}
// 		dtos, err = repo.NewMessageRepoInstance().MessageChatAll(
// 			&do.MessageDO{UserId: fromUserId, ToUserId: toUserId})
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	n := len(dtos)
// 	ret := make([]vo.Message, n)
// 	for i := 0; i < n; i++ {
// 		ret[i] = vo.Message{
// 			Id:         dtos[i].Id,
// 			FromUserId: dtos[i].UserId,
// 			ToUserId:   dtos[i].ToUserId,
// 			Content:    dtos[i].Content,
// 			CreateTime: dtos[i].CreateTime.Unix(),
// 		}
// 	}
// 	return ret, nil
// }

// func genChatKey(userIdA int64, userIdB int64) string {
// 	if userIdA > userIdB {
// 		return log.Sprintf("%d_%d", userIdB, userIdA)
// 	}
// 	return log.Sprintf("%d_%d", userIdA, userIdB)
// }

func MessageChat(dto *dto.MessageChatDTO) ([]vo.Message, error) {
	log.Println("##### MessageChat")
	fromUserId, err := util.JWTAuth(dto.Token)
	if err != nil {
		return nil, err
	}
	toUserId := dto.ToUserId
	log.Printf("##### %d %d from %d to %d\n", dto.PreMsgTime, time.Now().Unix(), fromUserId, toUserId)
	var dtos []do.MessageDO
	if dto.PreMsgTime == 0 {
		log.Println("##### 刚进聊天窗口")
		dtos, err = repo.NewMessageRepoInstance().MessageChatAll(
			&do.MessageDO{UserId: fromUserId, ToUserId: toUserId})
	} else {
		log.Println("##### 已在聊天窗口", util.UnixMilliToTime(dto.PreMsgTime))
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
	log.Printf("#####%#v", ret)
	return ret, nil
}
