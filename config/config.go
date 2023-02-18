package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type ConfigStruct struct {
	Mysql struct {
		Host     string
		Port     int
		Username string
		Password string
		Dbname   string
	}
	Minio struct {
		Host     string
		Port     int
		Username string
		Password string
	}
	StaticDir string
	// ChatPollingInterval int
}

var (
	Config ConfigStruct
)

const (
	FeedMaxNum = 5
)

func init() {
	log.Println("#####config.init()")
	workDir, _ := os.Getwd() //获取目录对应的路径
	LoadConfig(workDir + "/../../config")
}

func LoadConfig(fileDir string) {
	viper.SetConfigName("tiktok") //配置文件名
	viper.SetConfigType("yaml")   //配置文件类型
	viper.AddConfigPath(fileDir)  //执行go run对应的路径配置
	//viper.AddConfigPath(workDir+"/src/gin_application"+"/config") //执行单文件运行，
	log.Println(fileDir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("找不到配置文件..")
		} else {
			log.Println("配置文件出错..")
		}
	}

	allSettings := viper.AllSettings()
	log.Println(allSettings)

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Printf("Error unmarshalling config: %s", err)
	}
	log.Printf("%#v\n", Config)

	if err := os.MkdirAll(Config.StaticDir, 0755); err != nil {
		log.Println("创建静态资源目录失败..")
	}
}

// func LoadStruct() {
// 	// var config Config
// 	err := viper.Unmarshal(&Config)
// 	if err != nil {
// 		log.Printf("Error unmarshalling config: %s", err)
// 	}
// 	log.Println(Config)
// }
