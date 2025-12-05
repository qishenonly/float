package repository

import (
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/pkg/database"
)

// CategoryRepository 分类仓库接口
type CategoryRepository interface {
	Create(category *models.Category) error
	FindByID(id int64) (*models.Category, error)
	FindByUserID(userID int64, categoryType string) ([]*models.Category, error)
	FindAll(userID int64) ([]*models.Category, error)
	Update(category *models.Category) error
	Delete(id int64) error
	GetSystemCategories(categoryType string) ([]*models.Category, error)
}

type categoryRepository struct{}

// NewCategoryRepository 创建分类仓库实例
func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

// Create 创建分类
func (r *categoryRepository) Create(category *models.Category) error {
	return database.DB.Create(category).Error
}

// FindByID 根据ID查找分类
func (r *categoryRepository) FindByID(id int64) (*models.Category, error) {
	var category models.Category
	err := database.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// FindByUserID 根据用户ID和类型查找分类
func (r *categoryRepository) FindByUserID(userID int64, categoryType string) ([]*models.Category, error) {
	var categories []*models.Category
	query := database.DB.Where("user_id = ? AND is_active = ?", userID, true)

	if categoryType != "" {
		query = query.Where("type = ?", categoryType)
	}

	err := query.Order("display_order ASC, created_at ASC").Find(&categories).Error
	return categories, err
}

// FindAll 获取用户所有分类
func (r *categoryRepository) FindAll(userID int64) ([]*models.Category, error) {
	return r.FindByUserID(userID, "")
}

// Update 更新分类
func (r *categoryRepository) Update(category *models.Category) error {
	return database.DB.Save(category).Error
}

// Delete 删除分类
func (r *categoryRepository) Delete(id int64) error {
	return database.DB.Delete(&models.Category{}, id).Error
}

// GetSystemCategories 获取系统默认分类
func (r *categoryRepository) GetSystemCategories(categoryType string) ([]*models.Category, error) {
	var categories []*models.Category
	query := database.DB.Where("is_system = ? AND is_active = ?", true, true)

	if categoryType != "" {
		query = query.Where("type = ?", categoryType)
	}

	err := query.Order("display_order ASC").Find(&categories).Error
	return categories, err
}
