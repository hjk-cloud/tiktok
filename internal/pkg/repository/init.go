package repository

import (
	"fmt"
	"log"

	"github.com/hjk-cloud/tiktok/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// data source name
	dsn string
	Db  *gorm.DB
)

func setDsn() {
	host := config.Config.Mysql.Host
	port := config.Config.Mysql.Port
	user := config.Config.Mysql.Username
	pwd := config.Config.Mysql.Password
	dbname := config.Config.Mysql.Dbname
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", user, pwd, host, port, dbname)
}

// database
func init() {
	fmt.Println("#####repository.database.init()")
	setDsn()
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Panicln(err)
		// panic(err)
	}
	Db = db
}

// redis
func init() {
	fmt.Println("#####repository.redis.init()")
}

// func init() {

// }
