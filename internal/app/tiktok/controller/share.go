package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/vo"
)

func writeError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, vo.Response{
		StatusCode: 1,
		StatusMsg:  err.Error(),
	})
}
