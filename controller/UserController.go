package controller

import (
	"Gin_Vue_Demo/common"
	"Gin_Vue_Demo/dto"
	"Gin_Vue_Demo/model"
	"Gin_Vue_Demo/response"
	"Gin_Vue_Demo/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Info(c *gin.Context) {
	// 因为用户通过认证了，且在 AuthMiddleWare 中 c.Set("user",user) 将用户信息写入，所以可以直接拿
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			// 强转，user 是 interface{} 类型
			"user": dto.ToUserDto(user.(model.User)),
		},
	})
}

func Login(c *gin.Context) {
	db := common.GetDB()
	// 获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	// 判断手机号
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	// 判断密码
	// 加密密码判断，第一个是数据库加密过的密码，第二个是前端获取的密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	response.Success(c, http.StatusOK, 200, gin.H{"token": token}, "登录成功")
}
func Register(c *gin.Context) {
	db := common.GetDB()
	// 获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 数据验证
	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	// 如果没有传名称，给一个六位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(6)
	}

	// 判断手机号是否存在
	if isTelephoneExist(db, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "该手机号已经注册")
		return
	}

	// 创建用户
	// 加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// 系统级错误
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	db.Create(&newUser)
	// 返回结果
	response.Success(c, http.StatusOK, 200, nil, "注册成功")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
