package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/repository"
	"github.com/qiuhaonan/float-backend/internal/service"
	"github.com/qiuhaonan/float-backend/internal/utils"
	"github.com/qiuhaonan/float-backend/pkg/logger"
)

// TransactionHandler 交易处理器
type TransactionHandler struct {
	transactionService service.TransactionService
}

// NewTransactionHandler 创建交易处理器实例
func NewTransactionHandler() *TransactionHandler {
	return &TransactionHandler{
		transactionService: service.NewTransactionService(
			repository.NewTransactionRepository(),
			repository.NewAccountRepository(),
			repository.NewCategoryRepository(),
		),
	}
}

// CreateTransaction godoc
// @Summary 创建交易
// @Description 创建新的交易记录
// @Tags Transactions
// @Accept json
// @Produce json
// @Param body body request.CreateTransactionRequest true "交易信息"
// @Success 201 {object} response.TransactionResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/transactions [post]
// @Security Bearer
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		logger.Error("[Handler][创建交易] 未授权的请求")
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权的请求")
		return
	}

	var req request.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(fmt.Sprintf("[Handler][创建交易] 请求参数错误 | 用户ID: %d | 错误: %v", userID, err))
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][创建交易] 创建交易请求 | 用户ID: %d | 交易类型: %s | 金额: %v", userID, req.Type, req.Amount))

	resp, err := h.transactionService.CreateTransaction(userID, &req)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][创建交易] 创建失败 | 用户ID: %d | 错误: %v", userID, err))
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][创建交易] 创建成功 | 用户ID: %d | 交易ID: %d", userID, resp.ID))
	c.JSON(http.StatusCreated, utils.Response{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    resp,
	})
}

// CreateBatchTransactions godoc
// @Summary 批量创建交易
// @Description 批量创建多条交易记录
// @Tags Transactions
// @Accept json
// @Produce json
// @Param body body request.BulkCreateTransactionRequest true "交易列表"
// @Success 201 {object} response.BulkOperationResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/transactions/batch [post]
// @Security Bearer
func (h *TransactionHandler) CreateBatchTransactions(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		logger.Error("[Handler][批量创建交易] 未授权的请求")
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权的请求")
		return
	}

	var req request.BulkCreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(fmt.Sprintf("[Handler][批量创建交易] 请求参数错误 | 用户ID: %d | 错误: %v", userID, err))
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][批量创建交易] 批量创建请求 | 用户ID: %d | 数量: %d", userID, len(req.Transactions)))

	resp, err := h.transactionService.CreateBatchTransactions(userID, &req)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][批量创建交易] 创建失败 | 用户ID: %d | 错误: %v", userID, err))
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][批量创建交易] 创建成功 | 用户ID: %d | 成功数: %d | 失败数: %d", userID, resp.SuccessCount, resp.FailureCount))
	c.JSON(http.StatusCreated, utils.Response{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    resp,
	})
}

// GetTransaction godoc
// @Summary 获取交易详情
// @Description 根据ID获取交易详细信息
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path int true "交易ID"
// @Success 200 {object} response.TransactionResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/transactions/{id} [get]
// @Security Bearer
func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		logger.Error("[Handler][获取交易] 未授权的请求")
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权的请求")
		return
	}

	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][获取交易] 交易ID格式错误 | 用户ID: %d", userID))
		utils.ErrorResponse(c, http.StatusBadRequest, "交易ID格式错误")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][获取交易] 获取交易详情请求 | 用户ID: %d | 交易ID: %d", userID, transactionID))

	resp, err := h.transactionService.GetTransactionByID(userID, transactionID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][获取交易] 获取失败 | 用户ID: %d | 交易ID: %d | 错误: %v", userID, transactionID, err))
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][获取交易] 获取成功 | 用户ID: %d | 交易ID: %d", userID, transactionID))
	utils.SuccessResponse(c, resp)
}

// ListTransactions godoc
// @Summary 查询交易列表
// @Description 查询用户的交易列表，支持多种筛选条件
// @Tags Transactions
// @Accept json
// @Produce json
// @Param type query string false "交易类型 (expense/income/transfer)"
// @Param category_id query int false "分类ID"
// @Param account_id query int false "账户ID"
// @Param start_date query string false "开始日期 (YYYY-MM-DD)"
// @Param end_date query string false "结束日期 (YYYY-MM-DD)"
// @Param search_keyword query string false "搜索关键词"
// @Param sort_by query string false "排序字段 (date/amount)"
// @Param sort_order query string false "排序顺序 (asc/desc)"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} response.TransactionListResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/transactions [get]
// @Security Bearer
func (h *TransactionHandler) ListTransactions(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		logger.Error("[Handler][交易列表] 未授权的请求")
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权的请求")
		return
	}
	logger.Info(fmt.Sprintf("[Handler][交易列表] 查询交易列表请求 | 用户ID: %d", userID))

	var filters request.ListTransactionRequest

	// 解析查询参数
	filters.Type = c.Query("type")
	filters.SearchKeyword = c.Query("search_keyword")
	filters.SortBy = c.Query("sort_by")
	filters.SortOrder = c.Query("sort_order")

	// 解析数字参数
	if categoryID, err := strconv.ParseInt(c.Query("category_id"), 10, 64); err == nil && categoryID > 0 {
		filters.CategoryID = &categoryID
	}

	if accountID, err := strconv.ParseInt(c.Query("account_id"), 10, 64); err == nil && accountID > 0 {
		filters.AccountID = &accountID
	}

	if page, err := strconv.Atoi(c.Query("page")); err == nil && page > 0 {
		filters.Page = page
	} else {
		filters.Page = 1
	}

	if pageSize, err := strconv.Atoi(c.Query("page_size")); err == nil && pageSize > 0 {
		filters.PageSize = pageSize
	} else {
		filters.PageSize = 20
	}

	// 解析日期参数
	if startDateStr := c.Query("start_date"); startDateStr != "" {
		if startDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			filters.StartDate = startDate
		}
	}

	if endDateStr := c.Query("end_date"); endDateStr != "" {
		if endDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			filters.EndDate = endDate
		}
	}

	resp, err := h.transactionService.ListTransactions(userID, &filters)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][交易列表] 查询失败 | 用户ID: %d | 错误: %v", userID, err))
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][交易列表] 查询成功 | 用户ID: %d | 总数: %d", userID, resp.Total))
	utils.SuccessResponse(c, resp)
}

// UpdateTransaction godoc
// @Summary 更新交易
// @Description 更新已存在的交易记录
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path int true "交易ID"
// @Param body body request.UpdateTransactionRequest true "更新信息"
// @Success 200 {object} response.TransactionResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/transactions/{id} [put]
// @Security Bearer
func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		logger.Error("[Handler][更新交易] 未授权的请求")
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权的请求")
		return
	}

	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][更新交易] 交易ID格式错误 | 用户ID: %d", userID))
		utils.ErrorResponse(c, http.StatusBadRequest, "交易ID格式错误")
		return
	}

	var req request.UpdateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(fmt.Sprintf("[Handler][更新交易] 请求参数错误 | 用户ID: %d | 交易ID: %d | 错误: %v", userID, transactionID, err))
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][更新交易] 更新交易请求 | 用户ID: %d | 交易ID: %d", userID, transactionID))

	resp, err := h.transactionService.UpdateTransaction(userID, transactionID, &req)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][更新交易] 更新失败 | 用户ID: %d | 交易ID: %d | 错误: %v", userID, transactionID, err))
		if err.Error() == "transaction not found" {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		}
		return
	}

	logger.Info(fmt.Sprintf("[Handler][更新交易] 更新成功 | 用户ID: %d | 交易ID: %d", userID, transactionID))
	utils.SuccessResponse(c, resp)
}

// DeleteTransaction godoc
// @Summary 删除交易
// @Description 删除指定的交易记录
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path int true "交易ID"
// @Success 204
// @Failure 404 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/transactions/{id} [delete]
// @Security Bearer
func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		logger.Error("[Handler][删除交易] 未授权的请求")
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权的请求")
		return
	}

	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][删除交易] 交易ID格式错误 | 用户ID: %d", userID))
		utils.ErrorResponse(c, http.StatusBadRequest, "交易ID格式错误")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][删除交易] 删除交易请求 | 用户ID: %d | 交易ID: %d", userID, transactionID))

	if err := h.transactionService.DeleteTransaction(userID, transactionID); err != nil {
		logger.Error(fmt.Sprintf("[Handler][删除交易] 删除失败 | 用户ID: %d | 交易ID: %d | 错误: %v", userID, transactionID, err))
		if err.Error() == "transaction not found" {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	logger.Info(fmt.Sprintf("[Handler][删除交易] 删除成功 | 用户ID: %d | 交易ID: %d", userID, transactionID))
	c.JSON(http.StatusNoContent, nil)
}

// DeleteBatchTransactions godoc
// @Summary 批量删除交易
// @Description 批量删除多条交易记录
// @Tags Transactions
// @Accept json
// @Produce json
// @Param body body request.BulkDeleteTransactionRequest true "交易ID列表"
// @Success 200 {object} response.BulkOperationResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/transactions/batch [delete]
// @Security Bearer
func (h *TransactionHandler) DeleteBatchTransactions(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权的请求")
		return
	}

	var req request.BulkDeleteTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	resp, err := h.transactionService.DeleteBatchTransactions(userID, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, resp)
}

// GetTransactionStatistics godoc
// @Summary 获取交易统计
// @Description 获取指定日期范围内的交易统计信息
// @Tags Transactions
// @Accept json
// @Produce json
// @Param start_date query string false "开始日期 (YYYY-MM-DD)"
// @Param end_date query string false "结束日期 (YYYY-MM-DD)"
// @Success 200 {object} response.TransactionStatisticsResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/transactions/statistics [get]
// @Security Bearer
func (h *TransactionHandler) GetTransactionStatistics(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权的请求")
		return
	}

	startDate := time.Now().AddDate(0, -1, 0)
	if startDateStr := c.Query("start_date"); startDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = parsedDate
		}
	}

	endDate := time.Now()
	if endDateStr := c.Query("end_date"); endDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = parsedDate
		}
	}

	resp, err := h.transactionService.GetTransactionStatistics(userID, startDate, endDate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, resp)
}

// GetMonthlyStatistics godoc
// @Summary 获取月度统计
// @Description 获取指定月份的交易统计
// @Tags Transactions
// @Accept json
// @Produce json
// @Param month query string false "月份 (YYYY-MM)" default("2025-12")
// @Success 200 {object} []response.MonthlyStatisticsResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/transactions/monthly-statistics [get]
// @Security Bearer
func (h *TransactionHandler) GetMonthlyStatistics(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权的请求")
		return
	}

	monthStr := c.Query("month")
	if monthStr == "" {
		monthStr = time.Now().Format("2006-01")
	}

	month, err := time.Parse("2006-01", monthStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "月份格式错误，应为 YYYY-MM")
		return
	}

	resp, err := h.transactionService.GetMonthlyStatistics(userID, month)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, resp)
}

// GetCategoryStatistics godoc
// @Summary 获取分类统计
// @Description 获取指定日期范围内的分类统计信息
// @Tags Transactions
// @Accept json
// @Produce json
// @Param start_date query string false "开始日期 (YYYY-MM-DD)"
// @Param end_date query string false "结束日期 (YYYY-MM-DD)"
// @Success 200 {object} []response.CategoryStatisticsResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/transactions/category-statistics [get]
// @Security Bearer
func (h *TransactionHandler) GetCategoryStatistics(c *gin.Context) {
	userID := c.GetInt64("user_id")
	if userID == 0 {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权的请求")
		return
	}

	startDate := time.Now().AddDate(0, -1, 0)
	if startDateStr := c.Query("start_date"); startDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = parsedDate
		}
	}

	endDate := time.Now()
	if endDateStr := c.Query("end_date"); endDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = parsedDate
		}
	}

	resp, err := h.transactionService.GetCategoryStatistics(userID, startDate, endDate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, resp)
}
