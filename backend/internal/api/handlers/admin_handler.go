package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/service"
	"github.com/qiuhaonan/float-backend/internal/utils"
	"github.com/qiuhaonan/float-backend/pkg/logger"
)

type AdminHandler struct {
	userService service.UserService
}

func NewAdminHandler(userService service.UserService) *AdminHandler {
	return &AdminHandler{
		userService: userService,
	}
}

// GetSystemOverview 获取系统概览
func (h *AdminHandler) GetSystemOverview(c *gin.Context) {
	logger.Info("[Admin] 获取系统概览")

	overview, err := h.userService.GetSystemOverview()
	if err != nil {
		logger.Error(fmt.Sprintf("[Admin] 获取概览失败: %v", err))
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, overview)
}

// ListUsers 获取用户列表
func (h *AdminHandler) ListUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "20")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	logger.Info(fmt.Sprintf("[Admin] 获取用户列表 | Page: %d | Size: %d", page, pageSize))

	users, total, err := h.userService.ListUsers(page, pageSize)
	if err != nil {
		logger.Error(fmt.Sprintf("[Admin] 获取用户列表失败: %v", err))
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{
		"items": users,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}
