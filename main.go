package main

import (
	"Gin_Vue_Demo/common"
	"Gin_Vue_Demo/routers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	common.InitDB()

	r := gin.Default()

	r = routers.CollectRoute(r)

	panic(r.Run(":9090"))
}
