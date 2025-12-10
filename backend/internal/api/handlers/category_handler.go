package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/service"
	"github.com/qiuhaonan/float-backend/internal/utils"
	"github.com/qiuhaonan/float-backend/pkg/logger"
)

// CategoryHandler 分类处理器
type CategoryHandler struct {
	categoryService service.CategoryService
}

// NewCategoryHandler 创建分类处理器实例
func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		categoryService: service.NewCategoryService(),
	}
}

// GetCategories 获取分类列表
// @Summary 获取分类列表
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param type query string false "分类类型: expense/income"
// @Success 200 {object} utils.Response{data=[]response.CategoryResponse}
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		logger.Error("[Handler][分类列表] 用户未登录")
		utils.ErrorResponse(c, http.StatusUnauthorized, "未登录")
		return
	}

	uID := userID.(int64)
	// 获取查询参数
	categoryType := c.Query("type") // expense 或 income 或 空（全部）

	logger.Info(fmt.Sprintf("[Handler][分类列表] 获取分类列表请求 | 用户ID: %d | 分类类型: %s", uID, categoryType))

	// 获取分类列表
	categories, err := h.categoryService.GetCategories(uID, categoryType)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][分类列表] 获取失败 | 用户ID: %d | 错误: %v", uID, err))
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取分类失败")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][分类列表] 获取成功 | 用户ID: %d | 分类数: %d", uID, len(categories)))
	utils.SuccessResponse(c, categories)
}

// GetCategory 获取分类详情
// @Summary 获取分类详情
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param id path int true "分类ID"
// @Success 200 {object} utils.Response{data=response.CategoryResponse}
// @Router /categories/:id [get]
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uID := userID.(int64)

	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][分类详情] 分类ID格式错误 | 用户ID: %d", uID))
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的分类ID")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][分类详情] 获取分类详情请求 | 用户ID: %d | 分类ID: %d", uID, categoryID))

	category, err := h.categoryService.GetCategoryByID(uID, categoryID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][分类详情] 获取失败 | 用户ID: %d | 分类ID: %d | 错误: %v", uID, categoryID, err))
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][分类详情] 获取成功 | 用户ID: %d | 分类ID: %d", uID, categoryID))
	utils.SuccessResponse(c, category)
}

// CreateCategory 创建分类
// @Summary 创建自定义分类
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param category body request.CreateCategoryRequest true "分类信息"
// @Success 200 {object} utils.Response{data=response.CategoryResponse}
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uID := userID.(int64)

	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(fmt.Sprintf("[Handler][创建分类] 请求参数错误 | 用户ID: %d | 错误: %v", uID, err))
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][创建分类] 创建分类请求 | 用户ID: %d | 分类名: %s | 类型: %s", uID, req.Name, req.Type))

	category, err := h.categoryService.CreateCategory(uID, &req)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][创建分类] 创建失败 | 用户ID: %d | 错误: %v", uID, err))
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][创建分类] 创建成功 | 用户ID: %d | 分类ID: %d", uID, category.ID))
	utils.SuccessResponse(c, category)
}

// UpdateCategory 更新分类
// @Summary 更新分类
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param id path int true "分类ID"
// @Param category body request.UpdateCategoryRequest true "更新信息"
// @Success 200 {object} utils.Response
// @Router /categories/:id [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uID := userID.(int64)

	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][更新分类] 分类ID格式错误 | 用户ID: %d", uID))
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的分类ID")
		return
	}

	var req request.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(fmt.Sprintf("[Handler][更新分类] 请求参数错误 | 用户ID: %d | 分类ID: %d | 错误: %v", uID, categoryID, err))
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][更新分类] 更新分类请求 | 用户ID: %d | 分类ID: %d", uID, categoryID))

	if err := h.categoryService.UpdateCategory(uID, categoryID, &req); err != nil {
		logger.Error(fmt.Sprintf("[Handler][更新分类] 更新失败 | 用户ID: %d | 分类ID: %d | 错误: %v", uID, categoryID, err))
		utils.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][更新分类] 更新成功 | 用户ID: %d | 分类ID: %d", uID, categoryID))
	utils.SuccessResponse(c, nil)
}

// DeleteCategory 删除分类
// @Summary 删除分类
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param id path int true "分类ID"
// @Success 200 {object} utils.Response
// @Router /categories/:id [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uID := userID.(int64)

	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][删除分类] 分类ID格式错误 | 用户ID: %d", uID))
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的分类ID")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][删除分类] 删除分类请求 | 用户ID: %d | 分类ID: %d", uID, categoryID))

	if err := h.categoryService.DeleteCategory(uID, categoryID); err != nil {
		logger.Error(fmt.Sprintf("[Handler][删除分类] 删除失败 | 用户ID: %d | 分类ID: %d | 错误: %v", uID, categoryID, err))
		utils.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][删除分类] 删除成功 | 用户ID: %d | 分类ID: %d", uID, categoryID))
	utils.SuccessResponse(c, nil)
}

// GetSystemCategories 获取系统默认分类
// @Summary 获取系统默认分类
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param type query string false "分类类型: expense/income"
// @Success 200 {object} utils.Response{data=[]response.CategoryResponse}
// @Router /categories/system [get]
func (h *CategoryHandler) GetSystemCategories(c *gin.Context) {
	categoryType := c.Query("type")

	categories, err := h.categoryService.GetSystemCategories(categoryType)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取分类失败")
		return
	}

	utils.SuccessResponse(c, categories)
}
