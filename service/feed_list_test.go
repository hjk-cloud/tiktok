package service

import (
	"github.com/hjk-cloud/tiktok/model"
	"log"
	"testing"
)

func TestQueryFeedList(t *testing.T) {
	err := model.DBInit()
	if err != nil {
		return
	}
	//video, err := QueryFeedList(time.Now())
	if err != nil {
		log.Print(err)
	}
	//fmt.Println(video)
}
