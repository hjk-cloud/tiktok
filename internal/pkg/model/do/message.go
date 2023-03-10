package do

import (
	"time"
)

type MessageDO struct {
	Id         int64 `gorm:"primaryKey"`
	UserId     int64 `gorm:"index:user_from_to"` // 联合索引
	ToUserId   int64 `gorm:"index:user_from_to"` // 联合索引
	Content    string
	IsDeleted  bool
	CreateTime time.Time
	UpdateTime time.Time
}

func (MessageDO) TableName() string {
	return "t_message"
}
