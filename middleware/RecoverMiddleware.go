package middleware

import (
	"Gin_Vue_Demo/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
拦截错误的中间件
*/

func RecoveryMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if error := recover(); error != nil {
				response.Fail(c, nil, fmt.Sprint(error))
			}
		}()

		c.Next()
	}
}
