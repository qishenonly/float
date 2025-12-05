package response

import "time"

// CategoryResponse 分类响应
type CategoryResponse struct {
	ID           int64     `json:"id"`
	Type         string    `json:"type"`
	Name         string    `json:"name"`
	Icon         string    `json:"icon"`
	Color        string    `json:"color"`
	DisplayOrder int       `json:"display_order"`
	IsSystem     bool      `json:"is_system"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
