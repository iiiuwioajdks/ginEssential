package vo

type CreatedCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
