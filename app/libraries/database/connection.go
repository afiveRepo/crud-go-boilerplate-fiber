package database

import (
	"crud-go-boilerplate-fiber/app/models/entities"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMYSQL() *gorm.DB {
	// Connecting to MYSQL database
	config := map[string]string{
		"user":     viper.GetString("database.user"),
		"password": viper.GetString("database.pwd"),
		"host":     viper.GetString("database.host"),
		"port":     viper.GetString("database.port"),
		"db_name":  viper.GetString("database.db_name"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["user"],
		config["password"],
		config["host"],
		config["port"],
		config["db_name"],
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connection database")
	}
	db.AutoMigrate(entities.Product{}, entities.User{})
	return db
}
