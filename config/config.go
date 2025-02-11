package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct { //首字母一定要大写，否则无法读取，这是Unmarshal的规则，名称也要和yml文件的内容相匹配
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		log.Fatalf("Fatal error config file: %v", err)
	}

	AppConfig = &Config{}
	err = viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	initDB()
}
