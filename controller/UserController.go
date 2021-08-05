package controller

import (
	"Gin_Vue_Demo/common"
	"Gin_Vue_Demo/model"
	"Gin_Vue_Demo/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	// 获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能小于6位",
		})
		return
	}
	// 如果没有传名称，给一个六位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(6)
		return
	}
	log.Println(telephone, name, password)

	// 判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "该手机号已经注册",
		})
		return
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
