package main

import (
	"Gin_Vue_Demo/common"
	"Gin_Vue_Demo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()

	r := gin.Default()

	r = routers.CollectRoute(r)

	panic(r.Run(":9090"))
}
