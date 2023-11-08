package database

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("env_development")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./app/libraries/config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Fatal Error to Read Config \n", err)
		os.Exit(1)
	}
}
