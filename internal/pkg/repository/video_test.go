package repository

import (
	"fmt"
	_ "github.com/hjk-cloud/tiktok/config"
	"log"
	"testing"
)

// X 无法读取配置文件
func TestVideoRepo_QueryVideoById(t *testing.T) {

	video, err := videoRepo.QueryVideoById(1)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(video)
}
