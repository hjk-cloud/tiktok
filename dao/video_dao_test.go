package dao

import (
	"fmt"
	"github.com/hjk-cloud/tiktok/model"
	"log"
	"os"
	"testing"
	"time"
)

func TestVideoDao_QueryVideoById(t *testing.T) {
	err := model.DBInit()
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}
	video, err := videoDao.QueryVideoById(1)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(video)

}

func TestVideoDao_MQueryVideoByLastTime(t *testing.T) {
	err := model.DBInit()
	if err != nil {
		fmt.Println(err)
	}
	videos, err := videoDao.MQueryVideoByLastTime(time.Now())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(videos)
}
