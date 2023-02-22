package do

import (
	"time"

	"gorm.io/gorm"
)

type VideoDO struct {
	// gorm.Model
	Id            int64      `gorm:"primaryKey"`
	AuthorId      int64      `gorm:"index:idx_user_id;"`
	Title         string     // 自动映射成title
	PlayUrl       string     // 自动映射成play_url
	CoverUrl      string     // 自动映射成cover_url
	FavoriteCount int64      // 自动映射成favorite_count
	CommentCount  int64      // 自动映射成comment_count
	Status        byte       // 自动映射成status
	HashValue     string     `gorm:"column:hash;"` // 自动映射成hash_value，不一致，所以要手动指定
	CreateTime    time.Time  // 自动映射成create_time
	UpdateTime    *time.Time // 指针可以是nil，对应可空字段
}

func (VideoDO) TableName() string {
	return "t_video" // 自动映射成`videos`，不一致，所以要这样指定
}

// 注意，在 GORM 中保存、删除操作会【默认】运行在事务上。
// 因此在事务完成之前该事务中所作的更改是不可见的，如果您的钩子返回了任何错误，则修改将被回滚。
// 投稿后作品数+1
func (v *VideoDO) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&UserInfo{Id: v.AuthorId}).Update("publish_count", gorm.Expr("publish_count + 1")).Error
}
