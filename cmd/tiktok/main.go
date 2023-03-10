package main

// 包名改成tiktok后，go build会有问题

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hjk-cloud/tiktok/config"
	_ "github.com/hjk-cloud/tiktok/internal/pkg/repository"
)

// config 要在 repository 前 init

func main() {
	// go service.RunMessageServer()	// 另一种消息实现方案，未实现
	r := gin.Default()

	initRouter(r)
	// config.LoadStruct()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
