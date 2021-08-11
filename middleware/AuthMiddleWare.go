package middleware

import (
	"Gin_Vue_Demo/common"
	"Gin_Vue_Demo/model"
	"Gin_Vue_Demo/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 中间件

func AuthMiddleWare() gin.HandlerFunc {
	//
	return func(c *gin.Context) {
		// 获取 authorization
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		// 如果 error 或者 token 无效，返回权限不足
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		// 验证通过后获取 claim 中的 userId，也就是用户数据库里面的 id
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		// 用户存在,将用户信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
