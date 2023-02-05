package model

import (
	"github.com/hjk-cloud/tiktok/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化并连接数据库
var Db *gorm.DB

func DBInit() error {
	dsn := configs.DbUsername + ":" +
		configs.DbPassword + "@tcp(" + configs.DbUrl + ":3306)/" +
		configs.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err
}
