package repository

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	connLink string
	Db       *gorm.DB
)

func getConfig() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.pwd")
	dbname := viper.GetString("mysql.dbname")
	connLink = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, pwd, host, port, dbname)
}

// database
func init() {
	fmt.Println("#####repository.database.init()")
	getConfig()
	db, err := gorm.Open(mysql.Open(connLink))
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
