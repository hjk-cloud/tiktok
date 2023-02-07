package repository

import (
	"fmt"
	"log"
	"sync"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/entity"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TableName why？
// type Video entity.Video

// type Video struct {
// 	entity.Video
// }

var (
	connLink string
	db       *sql.DB
	gdb      *gorm.DB
)

func getConfig() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.pwd")
	dbname := viper.GetString("mysql.dbname")
	connLink = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, pwd, host, port, dbname)
}

func NewGormDB() *gorm.DB {
	getConfig()
	log.Println("#####", connLink)
	db, err := gorm.Open(mysql.Open(connLink))
	if err != nil {
		log.Panicln(err)
		// panic(err)
	}
	return db
}

func NewDB() *sql.DB {
	getConfig()
	DB, err := sql.Open("mysql", connLink) // config.DB_LINK) //"root:root@tcp(127.0.0.1:3306)/tiktok"
	if err != nil {
		log.Println(err)
		// panic(err)
	}
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

func (*VideoRepo) Create(video *(entity.Video)) (int64, error) {
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

// 奇怪
func (*VideoRepo) CreateByGorm(video *(entity.Video)) (int64, error) {
	fmt.Println("#####", video.TableName())
	gdb = NewGormDB()
	gdb.Create(video)
	return video.Id, nil
}
