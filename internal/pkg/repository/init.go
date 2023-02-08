package repository

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var mydb = NewGormDB()

func myGetConfig() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.pwd")
	dbname := viper.GetString("mysql.dbname")
	connLink = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, pwd, host, port, dbname)
}

func myNewGormDB() *gorm.DB {
	getConfig()
	log.Println("#####", connLink)
	db, err := gorm.Open(mysql.Open(connLink))
	if err != nil {
		log.Panicln(err)
		// panic(err)
	}
	return db
}

func myNewDB() *sql.DB {
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