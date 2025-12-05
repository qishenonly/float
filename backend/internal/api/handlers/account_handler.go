package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/service"
	"github.com/qiuhaonan/float-backend/internal/utils"
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

	accounts, err := h.accountService.GetAccounts(userID.(int64))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取账户失败")
		return
	}

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

	accountID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的账户ID")
		return
	}

	account, err := h.accountService.GetAccountByID(userID.(int64), accountID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

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

	var req request.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	account, err := h.accountService.CreateAccount(userID.(int64), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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

	accountID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的账户ID")
		return
	}

	var req request.UpdateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	if err := h.accountService.UpdateAccount(userID.(int64), accountID, &req); err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

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

	accountID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的账户ID")
		return
	}

	if err := h.accountService.DeleteAccount(userID.(int64), accountID); err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

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

	balance, err := h.accountService.GetAccountBalance(userID.(int64))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取余额失败")
		return
	}

	utils.SuccessResponse(c, balance)
}
