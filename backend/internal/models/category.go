package models

import (
	"time"
)

// Category 分类表
type Category struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int64     `gorm:"not null;index:idx_user_type" json:"user_id"`
	Type         string    `gorm:"type:enum('expense','income');not null;index:idx_user_type" json:"type"`
	Name         string    `gorm:"size:50;not null" json:"name"`
	Icon         string    `gorm:"size:50;not null" json:"icon"`
	Color        string    `gorm:"size:20;not null" json:"color"`
	DisplayOrder int       `gorm:"default:0" json:"display_order"`
	IsSystem     bool      `gorm:"default:false" json:"is_system"`
	IsActive     bool      `gorm:"default:true;index" json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName 表名
func (Category) TableName() string {
	return "categories"
}
