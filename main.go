package main

import (
	"Gin_Vue_Demo/common"
	"Gin_Vue_Demo/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	common.InitDB()

	r := gin.Default()

	r = routers.CollectRoute(r)

	port := viper.GetString("server.port")

	panic(r.Run(":" + port))
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
