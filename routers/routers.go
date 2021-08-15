package routers

import (
	"Gin_Vue_Demo/controller"
	"Gin_Vue_Demo/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleWare(), middleware.RecoveryMiddleWare())
	r.POST("/api/auth/register", controller.Register)

	r.POST("/api/auth/login", controller.Login)

	// 用中间件 token 认证保护用户信息接口, 并且将用户的信息传到前端
	r.GET("/api/auth/info", middleware.AuthMiddleWare(), controller.Info)

	category := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	category.POST("/", categoryController.Create)
	category.PUT("/:id", categoryController.Update)
	category.GET("/:id", categoryController.Show)
	category.DELETE("/:id", categoryController.Delete)

	post := r.Group("/posts")
	post.Use(middleware.AuthMiddleWare())
	postController := controller.NewPostController()
	post.POST("/", postController.Create)
	post.PUT("/:id", postController.Update)
	post.GET("/:id", postController.Show)
	post.DELETE("/:id", postController.Delete)
	post.POST("/page/list", postController.PageList)

	return r
}
