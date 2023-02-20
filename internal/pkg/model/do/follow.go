package do

import "time"

type Follow struct {
	SubjectId  int64
	ObjectId   int64
	isDeleted  int8
	CreateTime time.Time
	updateTime time.Time
}

func (Follow) TableName() string {
	return "t_follow"
}
