package flow

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type PublishActionFlow struct {
	Context *gin.Context
	Token   string
	Title   string
	Data    *multipart.FileHeader
	UserId  int64
	// Video  *Video
}
