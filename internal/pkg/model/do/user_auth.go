package do

import "time"

type UserAuth struct {
	Id         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Password   string `json:"password,omitempty"`
	isDeleted  int8
	CreateTime time.Time
	updateTime time.Time
}

func (UserAuth) TableName() string {
	return "t_user_auth"
}
