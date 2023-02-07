package dao

import (
	"fmt"
	"sync"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/entity"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
)

type Video entity.Video

var db *sql.DB

func NewDB() *sql.DB {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.pwd")
	dbname := viper.GetString("mysql.dbname")
	connLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, pwd, host, port, dbname)

	DB, _ := sql.Open("mysql", connLink) // config.DB_LINK) //"root:root@tcp(127.0.0.1:3306)/tiktok"
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}
	fmt.Println("connnect success")
	return DB
}

type VideoRepo struct {
}

var videoRepo *VideoRepo
var videoOnce sync.Once

func NewVideoRepoInstance() *VideoRepo {
	videoOnce.Do(
		func() {
			videoRepo = &VideoRepo{}
		})
	return videoRepo
}

func (*VideoRepo) Create(video *Video) (int64, error) {
	db = NewDB()

	sqlStr := "INSERT INTO t_video (`author_id`, `title`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `status`, `hash`, `create_time`)" +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	ret, err := db.Exec(sqlStr, video.AuthorId, video.Title, video.PlayUrl, video.CoverUrl, video.FavoriteCount,
		video.CommentCount, video.Status, video.HashValue, video.CreateTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return -1, err
	}
	newId, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return -1, err
	}
	fmt.Printf("insert success, the id is %d.\n", newId)

	return newId, nil
}
