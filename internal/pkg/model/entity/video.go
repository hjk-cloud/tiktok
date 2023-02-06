package entity

import (
	"time"
)

type Video struct {
	Id            int64
	AuthorId      int64
	Title         string
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int
	CommentCount  int
	Status        byte
	HashValue     string
	CreateTime    time.Time
	UpdateTime    time.Time
}
