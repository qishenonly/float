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

// AccountHandler 账户处理器
type AccountHandler struct {
	accountService service.AccountService
}

// NewAccountHandler 创建账户处理器实例
func NewAccountHandler() *AccountHandler {
	return &AccountHandler{
		accountService: service.NewAccountService(),
	}
}

// GetAccounts 获取账户列表
// @Summary 获取账户列表
// @Tags 账户管理
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response{data=[]response.AccountResponse}
// @Router /accounts [get]
func (h *AccountHandler) GetAccounts(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uID := userID.(int64)
	logger.Info(fmt.Sprintf("[Handler][账户列表] 获取账户列表请求 | 用户ID: %d", uID))

	accounts, err := h.accountService.GetAccounts(uID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][账户列表] 获取失败 | 用户ID: %d | 错误: %v", uID, err))
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取账户失败")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][账户列表] 获取成功 | 用户ID: %d", uID))
	utils.SuccessResponse(c, accounts)
}

// GetAccount 获取账户详情
// @Summary 获取账户详情
// @Tags 账户管理
// @Accept json
// @Produce json
// @Param id path int true "账户ID"
// @Success 200 {object} utils.Response{data=response.AccountResponse}
// @Router /accounts/:id [get]
func (h *AccountHandler) GetAccount(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uID := userID.(int64)

	accountID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][账户详情] 账户ID格式错误 | 用户ID: %d", uID))
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的账户ID")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][账户详情] 获取账户详情请求 | 用户ID: %d | 账户ID: %d", uID, accountID))

	account, err := h.accountService.GetAccountByID(uID, accountID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][账户详情] 获取失败 | 用户ID: %d | 账户ID: %d | 错误: %v", uID, accountID, err))
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][账户详情] 获取成功 | 用户ID: %d | 账户ID: %d", uID, accountID))
	utils.SuccessResponse(c, account)
}

// CreateAccount 创建账户
// @Summary 创建账户
// @Tags 账户管理
// @Accept json
// @Produce json
// @Param account body request.CreateAccountRequest true "账户信息"
// @Success 200 {object} utils.Response{data=response.AccountResponse}
// @Router /accounts [post]
func (h *AccountHandler) CreateAccount(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uID := userID.(int64)

	var req request.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(fmt.Sprintf("[Handler][创建账户] 请求参数错误 | 用户ID: %d | 错误: %v", uID, err))
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][创建账户] 创建账户请求 | 用户ID: %d | 账户名: %s", uID, req.AccountName))

	account, err := h.accountService.CreateAccount(uID, &req)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][创建账户] 创建失败 | 用户ID: %d | 错误: %v", uID, err))
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][创建账户] 创建成功 | 用户ID: %d | 账户ID: %d", uID, account.ID))
	utils.SuccessResponse(c, account)
}

// UpdateAccount 更新账户
// @Summary 更新账户
// @Tags 账户管理
// @Accept json
// @Produce json
// @Param id path int true "账户ID"
// @Param account body request.UpdateAccountRequest true "更新信息"
// @Success 200 {object} utils.Response
// @Router /accounts/:id [put]
func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uID := userID.(int64)

	accountID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][更新账户] 账户ID格式错误 | 用户ID: %d", uID))
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的账户ID")
		return
	}

	var req request.UpdateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(fmt.Sprintf("[Handler][更新账户] 请求参数错误 | 用户ID: %d | 账户ID: %d | 错误: %v", uID, accountID, err))
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][更新账户] 更新账户请求 | 用户ID: %d | 账户ID: %d", uID, accountID))

	if err := h.accountService.UpdateAccount(uID, accountID, &req); err != nil {
		logger.Error(fmt.Sprintf("[Handler][更新账户] 更新失败 | 用户ID: %d | 账户ID: %d | 错误: %v", uID, accountID, err))
		utils.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][更新账户] 更新成功 | 用户ID: %d | 账户ID: %d", uID, accountID))
	utils.SuccessResponse(c, nil)
}

// DeleteAccount 删除账户
// @Summary 删除账户
// @Tags 账户管理
// @Accept json
// @Produce json
// @Param id path int true "账户ID"
// @Success 200 {object} utils.Response
// @Router /accounts/:id [delete]
func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uID := userID.(int64)

	accountID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][删除账户] 账户ID格式错误 | 用户ID: %d", uID))
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的账户ID")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][删除账户] 删除账户请求 | 用户ID: %d | 账户ID: %d", uID, accountID))

	if err := h.accountService.DeleteAccount(uID, accountID); err != nil {
		logger.Error(fmt.Sprintf("[Handler][删除账户] 删除失败 | 用户ID: %d | 账户ID: %d | 错误: %v", uID, accountID, err))
		utils.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	logger.Info(fmt.Sprintf("[Handler][删除账户] 删除成功 | 用户ID: %d | 账户ID: %d", uID, accountID))
	utils.SuccessResponse(c, nil)
}

// GetAccountBalance 获取账户余额汇总
// @Summary 获取账户余额汇总
// @Tags 账户管理
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response{data=response.AccountBalanceResponse}
// @Router /accounts/balance [get]
func (h *AccountHandler) GetAccountBalance(c *gin.Context) {
	userID, _ := c.Get("user_id")
	uID := userID.(int64)
	logger.Info(fmt.Sprintf("[Handler][账户余额] 获取账户余额请求 | 用户ID: %d", uID))

	balance, err := h.accountService.GetAccountBalance(uID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][账户余额] 获取失败 | 用户ID: %d | 错误: %v", uID, err))
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取余额失败")
		return
	}

	logger.Info(fmt.Sprintf("[Handler][账户余额] 获取成功 | 用户ID: %d", uID))
	utils.SuccessResponse(c, balance)
}
