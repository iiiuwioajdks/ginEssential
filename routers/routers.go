package routers

import (
	"Gin_Vue_Demo/controller"
	"Gin_Vue_Demo/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)

	r.POST("/api/auth/login", controller.Login)

	// 用中间件 token 认证保护用户信息接口, 并且将用户的信息传到前端
	r.GET("/api/auth/info", middleware.AuthMiddleWare(), controller.Info)

	return r
}
