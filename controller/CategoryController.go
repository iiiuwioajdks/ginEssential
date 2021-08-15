package controller

import (
	"Gin_Vue_Demo/model"
	"Gin_Vue_Demo/repository"
	"Gin_Vue_Demo/response"
	"Gin_Vue_Demo/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
}

// alt+insert 可以继承方法

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})

	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreatedCategoryRequest

	//利用 gin 的绑定参数
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, gin.H{"err": err}, "数据验证错误，分类名称必填")
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory = model.Category{}
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	// 获取path中的 categoryId ，强转成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, gin.H{"err": err}, "分类不存在")
		return
	}

	// map , struct , name value
	updateCategory, err := c.Repository.Update(*category, requestCategory.Name)
	if err != nil {
		// panic 完交给 recoverMiddleWare 中间件去处理
		panic(err)
	}

	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的 categoryId ，强转成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "查询成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的 categoryId ，强转成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	err := c.Repository.DeleteById(categoryId)
	if err != nil {
		response.Fail(ctx, gin.H{"err": err}, "删除失败")
		return
	}

	response.Success(ctx, nil, "删除成功")

}
