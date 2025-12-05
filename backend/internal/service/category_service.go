package service

import (
	"errors"

	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/dto/response"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/internal/repository"
)

// CategoryService 分类服务接口
type CategoryService interface {
	GetCategories(userID int64, categoryType string) ([]*response.CategoryResponse, error)
	GetCategoryByID(userID int64, categoryID int64) (*response.CategoryResponse, error)
	CreateCategory(userID int64, req *request.CreateCategoryRequest) (*response.CategoryResponse, error)
	UpdateCategory(userID int64, categoryID int64, req *request.UpdateCategoryRequest) error
	DeleteCategory(userID int64, categoryID int64) error
	GetSystemCategories(categoryType string) ([]*response.CategoryResponse, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

// NewCategoryService 创建分类服务实例
func NewCategoryService() CategoryService {
	return &categoryService{
		categoryRepo: repository.NewCategoryRepository(),
	}
}

// GetCategories 获取分类列表（包含用户自定义+系统默认）
func (s *categoryService) GetCategories(userID int64, categoryType string) ([]*response.CategoryResponse, error) {
	// 获取用户自定义分类
	userCategories, err := s.categoryRepo.FindByUserID(userID, categoryType)
	if err != nil {
		return nil, err
	}

	// 获取系统默认分类
	systemCategories, err := s.categoryRepo.GetSystemCategories(categoryType)
	if err != nil {
		return nil, err
	}

	// 合并结果
	var result []*response.CategoryResponse

	// 添加用户自定义分类
	for _, cat := range userCategories {
		result = append(result, s.toCategoryResponse(cat))
	}

	// 添加系统分类
	for _, cat := range systemCategories {
		result = append(result, s.toCategoryResponse(cat))
	}

	return result, nil
}

// GetCategoryByID 获取分类详情
func (s *categoryService) GetCategoryByID(userID int64, categoryID int64) (*response.CategoryResponse, error) {
	category, err := s.categoryRepo.FindByID(categoryID)
	if err != nil {
		return nil, err
	}

	// 验证权限：只能查看自己的分类或系统分类
	if !category.IsSystem && category.UserID != userID {
		return nil, errors.New("无权访问该分类")
	}

	return s.toCategoryResponse(category), nil
}

// CreateCategory 创建分类
func (s *categoryService) CreateCategory(userID int64, req *request.CreateCategoryRequest) (*response.CategoryResponse, error) {
	// 创建分类对象
	category := &models.Category{
		UserID:       userID,
		Type:         req.Type,
		Name:         req.Name,
		Icon:         req.Icon,
		Color:        req.Color,
		DisplayOrder: req.DisplayOrder,
		IsSystem:     false,
		IsActive:     true,
	}

	if err := s.categoryRepo.Create(category); err != nil {
		return nil, errors.New("创建分类失败")
	}

	return s.toCategoryResponse(category), nil
}

// UpdateCategory 更新分类
func (s *categoryService) UpdateCategory(userID int64, categoryID int64, req *request.UpdateCategoryRequest) error {
	// 查找分类
	category, err := s.categoryRepo.FindByID(categoryID)
	if err != nil {
		return err
	}

	// 验证权限
	if category.IsSystem {
		return errors.New("系统分类不可修改")
	}
	if category.UserID != userID {
		return errors.New("无权修改该分类")
	}

	// 更新字段
	if req.Name != nil {
		category.Name = *req.Name
	}
	if req.Icon != nil {
		category.Icon = *req.Icon
	}
	if req.Color != nil {
		category.Color = *req.Color
	}
	if req.DisplayOrder != nil {
		category.DisplayOrder = *req.DisplayOrder
	}
	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}

	return s.categoryRepo.Update(category)
}

// DeleteCategory 删除分类
func (s *categoryService) DeleteCategory(userID int64, categoryID int64) error {
	// 查找分类
	category, err := s.categoryRepo.FindByID(categoryID)
	if err != nil {
		return err
	}

	// 验证权限
	if category.IsSystem {
		return errors.New("系统分类不可删除")
	}
	if category.UserID != userID {
		return errors.New("无权删除该分类")
	}

	return s.categoryRepo.Delete(categoryID)
}

// GetSystemCategories 获取系统默认分类
func (s *categoryService) GetSystemCategories(categoryType string) ([]*response.CategoryResponse, error) {
	categories, err := s.categoryRepo.GetSystemCategories(categoryType)
	if err != nil {
		return nil, err
	}

	var result []*response.CategoryResponse
	for _, cat := range categories {
		result = append(result, s.toCategoryResponse(cat))
	}

	return result, nil
}

// toCategoryResponse 转换为响应对象
func (s *categoryService) toCategoryResponse(category *models.Category) *response.CategoryResponse {
	return &response.CategoryResponse{
		ID:           category.ID,
		Type:         category.Type,
		Name:         category.Name,
		Icon:         category.Icon,
		Color:        category.Color,
		DisplayOrder: category.DisplayOrder,
		IsSystem:     category.IsSystem,
		IsActive:     category.IsActive,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}
}
