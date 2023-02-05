package model

import "time"

type Video struct {
	Id            int64
	AuthorId      int64
	Title         string
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int
	CommentCount  int
	Status        byte
	HashValue     string `gorm:"column:hash"`
	CreateTime    time.Time
	DeleteTime    time.Time
}

func (Video) TableName() string {
	return "t_video"
}
