package controller

import "github.com/gin-gonic/gin"

/**
增删改查的控制层接口，方便复用
*/

type RestController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Show(c *gin.Context)
	Delete(c *gin.Context)
}
