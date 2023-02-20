package dto

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type PublishActionDTO struct {
	Context *gin.Context
	Token   string
	Title   string
	Data    *multipart.FileHeader
}

type PublishListDTO struct {
	Token  string
	UserId int64
}
