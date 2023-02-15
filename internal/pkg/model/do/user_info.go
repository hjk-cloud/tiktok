package entity

import "time"

type UserInfo struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	TiktokId      string `json:"tiktok_id,omitempty"`
	Avatar        string `json:avatar,omitempty`
	Background    string `json:"background,omitempty"`
	Age           int8   `json:age,omitempty`
	Address       string `json:address,omitempty`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	isDelete      int8   `json:is_delete,omitempty`
	CreateTime    time.Time
	updateTime    time.Time
}

func (UserInfo) TableName() string {
	return "t_user_info"
}
