package do

import (
	"time"
)

type MessageDO struct {
	// gorm.Model
	Id        int64 `gorm:"primaryKey"`
	UserId    int64 // TODO 联合索引
	ToUserId  int64 // TODO 联合索引
	Content   string
	IsDeleted bool
	// IsRead     bool
	CreateTime time.Time
	UpdateTime time.Time
}

func (MessageDO) TableName() string {
	return "t_message"
}
