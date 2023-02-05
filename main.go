package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/model"
	"log"
	"os"
)

func main() {
	//go service.RunMessageServer()
	err := model.DBInit()
	if err != nil {
		log.Print("DBInit err: ", err)
		os.Exit(-1)
	}

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
