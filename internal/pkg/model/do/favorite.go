package do

import "time"

type Favorite struct {
	SubjectId  int64
	ObjectId   int64
	ObjectType string
	isDeleted  int8
	CreateTime time.Time
	updateTime time.Time
}

func (Favorite) TableName() string {
	return "t_favorite"
}
