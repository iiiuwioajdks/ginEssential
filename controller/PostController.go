package controller

import (
	"Gin_Vue_Demo/common"
	"Gin_Vue_Demo/model"
	"Gin_Vue_Demo/response"
	"Gin_Vue_Demo/vo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type IPostController interface {
	RestController
	PageList(c *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}

func (p PostController) Create(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := c.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "数据验证错误")
		return
	}

	// 获取登录用户,需要 AuthMiddleWare 中间件
	user, _ := c.Get("user")

	// 创建文章
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}
	response.Success(c, nil, "创建成功")
}

func (p PostController) Update(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := c.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "数据验证错误")
		return
	}
	// 获取path 中的id
	postId := c.Params.ByName("id")

	var post model.Post
	var count int64
	if p.DB.Where("id=?", postId).First(&post).Count(&count); count == 0 {
		response.Fail(c, nil, "文章不存在")
		return
	}
	// 判断当前用户是否为文章作者
	// 获取登录用户,需要 AuthMiddleWare 中间件
	user, _ := c.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(c, nil, "非法操作")
		return
	}

	// 更新文章
	toPost := changeToPost(&requestPost)
	if err := p.DB.Model(&post).Updates(toPost).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, gin.H{"post": post}, "更新成功")
}

func (p PostController) Show(c *gin.Context) {
	// 获取path 中的id
	postId := c.Params.ByName("id")

	var post model.Post
	var count int64
	if p.DB.Where("id=?", postId).First(&post).Count(&count); count == 0 {
		response.Fail(c, nil, "文章不存在")
		return
	}

	response.Success(c, gin.H{"post": post}, "成功")
}

func (p PostController) Delete(c *gin.Context) {
	// 获取path 中的id
	postId := c.Params.ByName("id")

	var post model.Post
	var count int64
	if p.DB.Where("id=?", postId).First(&post).Count(&count); count == 0 {
		response.Fail(c, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章作者
	// 获取登录用户,需要 AuthMiddleWare 中间件
	user, _ := c.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(c, nil, "非法操作")
		return
	}

	if err := p.DB.Delete(&post).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}

	response.Success(c, nil, "成功")
}

/**
给前端分页
*/

func (p PostController) PageList(c *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	// 分页
	var posts []model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 前端需要知道总条数
	var total int64
	p.DB.Model(model.Post{}).Count(&total)

	response.Success(c, gin.H{"data": posts, "total": total}, "成功")
}

func changeToPost(request *vo.CreatePostRequest) *model.Post {
	post := model.Post{
		CategoryId: request.CategoryId,
		Title:      request.Title,
		HeadImg:    request.HeadImg,
		Content:    request.Content,
	}
	return &post
}
