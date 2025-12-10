package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto_request "github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/service"
	"github.com/qiuhaonan/float-backend/internal/utils"
	"github.com/qiuhaonan/float-backend/pkg/cache"
	"github.com/qiuhaonan/float-backend/pkg/database"
)

// HealthCheck 健康检查
func HealthCheck(c *gin.Context) {
	// 检查数据库连接
	sqlDB, err := database.GetDB().DB()
	if err != nil || sqlDB.Ping() != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":   "unhealthy",
			"database": "down",
		})
		return
	}

	// 检查 Redis 连接
	if _, err := cache.GetClient().Ping(c.Request.Context()).Result(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"redis":  "down",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

// SendVerificationCode 发送邮箱验证码
func SendVerificationCode(c *gin.Context) {
	var req dto_request.SendVerificationCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	userService := service.NewUserService()
	if err := userService.SendVerificationCode(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "验证码已发送，请检查邮箱"})
}

// Register 用户注册
func Register(c *gin.Context) {
	var req dto_request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	userService := service.NewUserService()
	authResp, err := userService.Register(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, authResp)
}

// Login 用户登录
func Login(c *gin.Context) {
	var req dto_request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	userService := service.NewUserService()
	authResp, err := userService.Login(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SuccessResponse(c, authResp)
}

// RefreshToken 刷新Token
func RefreshToken(c *gin.Context) {
	var req dto_request.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	userService := service.NewUserService()
	tokenResp, err := userService.RefreshToken(req.RefreshToken)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SuccessResponse(c, tokenResp)
}

// GetCurrentUser 获取当前用户
func GetCurrentUser(c *gin.Context) {
	userID := c.GetInt64("user_id")

	userService := service.NewUserService()
	userResp, err := userService.GetUserByID(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, userResp)
}

// UpdateCurrentUser 更新当前用户
func UpdateCurrentUser(c *gin.Context) {
	userID := c.GetInt64("user_id")

	var req dto_request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	userService := service.NewUserService()
	if err := userService.UpdateUser(userID, &req); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "用户信息更新成功"})
}

// UpdatePassword 修改密码
func UpdatePassword(c *gin.Context) {
	userID := c.GetInt64("user_id")

	var req dto_request.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	userService := service.NewUserService()
	if err := userService.UpdatePassword(userID, &req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "密码修改成功"})
}

// GetUserStats 获取用户统计
func GetUserStats(c *gin.Context) {
	userID := c.GetInt64("user_id")

	userService := service.NewUserService()
	stats, err := userService.GetUserStats(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, stats)
}

// GetTransactions 获取交易列表
func GetTransactions(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetTransactions endpoint - to be implemented"})
}

// CreateTransaction 创建交易
func CreateTransaction(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "CreateTransaction endpoint - to be implemented"})
}

// GetTransaction 获取交易详情
func GetTransaction(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetTransaction endpoint - to be implemented"})
}

// UpdateTransaction 更新交易
func UpdateTransaction(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "UpdateTransaction endpoint - to be implemented"})
}

// DeleteTransaction 删除交易
func DeleteTransaction(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "DeleteTransaction endpoint - to be implemented"})
}

// GetTransactionStats 获取交易统计
func GetTransactionStats(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetTransactionStats endpoint - to be implemented"})
}

// GetAccounts 获取账户列表
func GetAccounts(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetAccounts endpoint - to be implemented"})
}

// CreateAccount 创建账户
func CreateAccount(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "CreateAccount endpoint - to be implemented"})
}

// GetAccount 获取账户详情
func GetAccount(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetAccount endpoint - to be implemented"})
}

// UpdateAccount 更新账户
func UpdateAccount(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "UpdateAccount endpoint - to be implemented"})
}

// DeleteAccount 删除账户
func DeleteAccount(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "DeleteAccount endpoint - to be implemented"})
}

// GetAccountBalance 获取账户余额
func GetAccountBalance(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetAccountBalance endpoint - to be implemented"})
}

// GetCategories 获取分类列表
func GetCategories(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetCategories endpoint - to be implemented"})
}

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "CreateCategory endpoint - to be implemented"})
}

// UpdateCategory 更新分类
func UpdateCategory(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "UpdateCategory endpoint - to be implemented"})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "DeleteCategory endpoint - to be implemented"})
}

// GetBills 获取账单列表
func GetBills(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetBills endpoint - to be implemented"})
}

// CreateBill 创建账单
func CreateBill(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "CreateBill endpoint - to be implemented"})
}

// GetBill 获取账单详情
func GetBill(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetBill endpoint - to be implemented"})
}

// UpdateBill 更新账单
func UpdateBill(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "UpdateBill endpoint - to be implemented"})
}

// DeleteBill 删除账单
func DeleteBill(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "DeleteBill endpoint - to be implemented"})
}

// GetUpcomingBills 获取即将到期账单
func GetUpcomingBills(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetUpcomingBills endpoint - to be implemented"})
}

// GetNotifications 获取通知列表
func GetNotifications(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetNotifications endpoint - to be implemented"})
}

// MarkNotificationRead 标记通知已读
func MarkNotificationRead(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "MarkNotificationRead endpoint - to be implemented"})
}

// GetUnreadCount 获取未读数量
func GetUnreadCount(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{"message": "GetUnreadCount endpoint - to be implemented"})
}
