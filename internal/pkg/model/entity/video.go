package entity

import (
	"time"
)

type Video struct {
	// gorm.Model
	Id            int64      `gorm:"primaryKey"`
	AuthorId      int64      `gorm:"index:idx_user_id;"`
	Title         string     // 自动映射成title
	PlayUrl       string     // 自动映射成play_url
	CoverUrl      string     // 自动映射成cover_url
	FavoriteCount int        // 自动映射成favorite_count
	CommentCount  int        // 自动映射成comment_count
	Status        byte       // 自动映射成status
	HashValue     string     `gorm:"column:hash;"` // 自动映射成hash_value，不一致，所以要手动指定
	CreateTime    time.Time  // 自动映射成create_time
	UpdateTime    *time.Time // 指针可以是nil，对应可空字段
}

func (Video) TableName() string {
	return "t_video" // 自动映射成`videos`，不一致，所以要这样指定
}
