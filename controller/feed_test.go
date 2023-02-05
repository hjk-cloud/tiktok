package controller

import (
	"github.com/hjk-cloud/tiktok/model"
	"testing"
)

func TestFeed(t *testing.T) {
	err := model.DBInit()
	if err != nil {
		return
	}
	Feed(nil)
}
