package repository

import (
	"log"
	"testing"

	_ "github.com/hjk-cloud/tiktok/config"
)

// X 无法读取配置文件
func TestVideoRepo_QueryVideoById(t *testing.T) {

	video, err := videoRepo.QueryVideoById(1)
	if err != nil {
		log.Print(err)
	}
	log.Println(video)
}
