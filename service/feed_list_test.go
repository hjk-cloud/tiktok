package service

import (
	"fmt"
	"github.com/hjk-cloud/tiktok/model"
	"log"
	"testing"
	"time"
)

func TestQueryFeedList(t *testing.T) {
	err := model.DBInit()
	if err != nil {
		return
	}
	video, err := QueryFeedList(time.Now())
	if err != nil {
		log.Print(err)
	}
	fmt.Println(video)
}
