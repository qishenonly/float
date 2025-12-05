package request

// CreateCategoryRequest 创建分类请求
type CreateCategoryRequest struct {
	Type         string `json:"type" binding:"required,oneof=expense income"`
	Name         string `json:"name" binding:"required,max=50"`
	Icon         string `json:"icon" binding:"required,max=50"`
	Color        string `json:"color" binding:"required,max=20"`
	DisplayOrder int    `json:"display_order"`
}

// UpdateCategoryRequest 更新分类请求
type UpdateCategoryRequest struct {
	Name         *string `json:"name" binding:"omitempty,max=50"`
	Icon         *string `json:"icon" binding:"omitempty,max=50"`
	Color        *string `json:"color" binding:"omitempty,max=20"`
	DisplayOrder *int    `json:"display_order"`
	IsActive     *bool   `json:"is_active"`
}
