package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	User struct {
		Name     string
		UseCsv   int
		Password string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}

	Csv struct {
		savaPath string
	}
}

var TodoConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.SetConfigFile("./config")

	TodoConfig = &Config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper load config.yml wrong!")
	}
	if err := viper.Unmarshal(TodoConfig); err != nil {
		log.Fatalf("init TodoConfig wrong!")
	}
}
