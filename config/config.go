package config

import (
	"fmt"
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
}

var (
	Config ConfigStruct
)

func init() {
	LoadConfig()
}

func LoadConfig() {
	fmt.Println("#####config.init()")
	workDir, _ := os.Getwd()                       //获取目录对应的路径
	viper.SetConfigName("tiktok")                  //配置文件名
	viper.SetConfigType("yaml")                    //配置文件类型
	viper.AddConfigPath(workDir + "/../../config") //执行go run对应的路径配置
	//viper.AddConfigPath(workDir+"/src/gin_application"+"/config") //执行单文件运行，
	fmt.Println(workDir+"../../config", workDir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	// host := viper.GetString("database.host")
	// fmt.Println("viper load yml: ", host)

	allSettings := viper.AllSettings()
	fmt.Println(allSettings)

	err := viper.Unmarshal(&Config)
	if err != nil {
		fmt.Printf("Error unmarshalling config: %s", err)
	}
	fmt.Printf("%#v\n", Config)
}

func LoadStruct() {
	// var config Config
	err := viper.Unmarshal(&Config)
	if err != nil {
		fmt.Printf("Error unmarshalling config: %s", err)
	}
	fmt.Println(Config)
}
