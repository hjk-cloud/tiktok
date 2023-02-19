package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
)

var chatConnMap = sync.Map{}

func RunMessageServer() {
	log.Println("##### RunMessageServer 9090")
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	defer listen.Close()
	if err != nil {
		log.Printf("Run message sever failed: %v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Accept conn failed: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	log.Println("##### Listen Accept process")

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break
			}
			log.Printf("Read message failed: %v\n", err)
			continue
		}

		// UserId 发送给 ToUserId
		var event = dto.MessageSendEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		log.Printf("Receive Message：%+v\n", event)

		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
		if len(event.MsgContent) == 0 {
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		// ToUserId 接收到 UserId
		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
		writeConn, exist := chatConnMap.Load(toChatKey)
		if !exist {
			log.Printf("User %d offline\n", event.ToUserId)
			continue
		}

		pushEvent := dto.MessagePushEvent{
			FromUserId: event.UserId,
			MsgContent: event.MsgContent,
		}
		pushData, _ := json.Marshal(pushEvent)
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			log.Printf("Push message failed: %v\n", err)
		}
	}
}
