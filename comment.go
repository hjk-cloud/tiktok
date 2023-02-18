package dao

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primaryKey"`
	UserId     int64  `json:"-"`
	VideoId    int64  `json:"video_id"`
	CreateDate string `json:"create_date,omitempty" gorm:"type:varchar(10);index"`
	Content    string `json:"content,omitempty" gorm:"type:text"`

	//调用接口
	User User `json:"user"`
}
