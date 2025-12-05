package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/service"
	"github.com/qiuhaonan/float-backend/internal/utils"
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
		utils.ErrorResponse(c, http.StatusUnauthorized, "未登录")
		return
	}

	// 获取查询参数
	categoryType := c.Query("type") // expense 或 income 或 空（全部）

	// 获取分类列表
	categories, err := h.categoryService.GetCategories(userID.(int64), categoryType)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取分类失败")
		return
	}

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

	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的分类ID")
		return
	}

	category, err := h.categoryService.GetCategoryByID(userID.(int64), categoryID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

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

	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	category, err := h.categoryService.CreateCategory(userID.(int64), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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

	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的分类ID")
		return
	}

	var req request.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	if err := h.categoryService.UpdateCategory(userID.(int64), categoryID, &req); err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

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

	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的分类ID")
		return
	}

	if err := h.categoryService.DeleteCategory(userID.(int64), categoryID); err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

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
