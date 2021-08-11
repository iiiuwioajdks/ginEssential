package common

import (
	"Gin_Vue_Demo/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password,
		host, port, database, charset)
	// "root:123456@(localhost:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
