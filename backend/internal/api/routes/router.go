package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/api/handlers"
	"github.com/qiuhaonan/float-backend/internal/api/middlewares"
	"github.com/qiuhaonan/float-backend/internal/repository"
	"github.com/qiuhaonan/float-backend/internal/service"
	"github.com/qiuhaonan/float-backend/pkg/database"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 中间件
	router.Use(middlewares.CORS())
	router.Use(middlewares.Logger())
	router.Use(middlewares.Recovery())

	// 健康检查
	router.GET("/health", handlers.HealthCheck)

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		// 认证路由（公开）
		auth := v1.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
			auth.POST("/refresh", handlers.RefreshToken)
		}

		// 需要认证的路由
		authorized := v1.Group("")
		authorized.Use(middlewares.AuthMiddleware())
		{
			// 用户管理
			users := authorized.Group("/users")
			{
				users.GET("/me", handlers.GetCurrentUser)
				users.PUT("/me", handlers.UpdateCurrentUser)
				users.PUT("/me/password", handlers.UpdatePassword)
				users.GET("/me/stats", handlers.GetUserStats)
			}

			// 交易记录
			transactions := authorized.Group("/transactions")
			{
				transactions.GET("", handlers.GetTransactions)
				transactions.POST("", handlers.CreateTransaction)
				transactions.GET("/:id", handlers.GetTransaction)
				transactions.PUT("/:id", handlers.UpdateTransaction)
				transactions.DELETE("/:id", handlers.DeleteTransaction)
				transactions.GET("/stats", handlers.GetTransactionStats)
			}

			// 账户管理
			accountHandler := handlers.NewAccountHandler()
			accounts := authorized.Group("/accounts")
			{
				accounts.GET("", accountHandler.GetAccounts)
				accounts.POST("", accountHandler.CreateAccount)
				accounts.GET("/balance", accountHandler.GetAccountBalance)
				accounts.GET("/:id", accountHandler.GetAccount)
				accounts.PUT("/:id", accountHandler.UpdateAccount)
				accounts.DELETE("/:id", accountHandler.DeleteAccount)
			}

			// 分类管理
			categoryHandler := handlers.NewCategoryHandler()
			categories := authorized.Group("/categories")
			{
				categories.GET("", categoryHandler.GetCategories)
				categories.POST("", categoryHandler.CreateCategory)
				categories.GET("/system", categoryHandler.GetSystemCategories)
				categories.GET("/:id", categoryHandler.GetCategory)
				categories.PUT("/:id", categoryHandler.UpdateCategory)
				categories.DELETE("/:id", categoryHandler.DeleteCategory)
			}

			// 账单订阅
			bills := authorized.Group("/bills")
			{
				bills.GET("", handlers.GetBills)
				bills.POST("", handlers.CreateBill)
				bills.GET("/:id", handlers.GetBill)
				bills.PUT("/:id", handlers.UpdateBill)
				bills.DELETE("/:id", handlers.DeleteBill)
				bills.GET("/upcoming", handlers.GetUpcomingBills)
			}

			// 通知
			notifications := authorized.Group("/notifications")
			{
				notifications.GET("", handlers.GetNotifications)
				notifications.PUT("/:id/read", handlers.MarkNotificationRead)
				notifications.GET("/unread", handlers.GetUnreadCount)
			}
		}

		// 软件更新 (公开接口或部分公开)
		appUpdateRepo := repository.NewAppUpdateRepository(database.DB)
		appUpdateService := service.NewAppUpdateService(appUpdateRepo)
		appUpdateHandler := handlers.NewAppUpdateHandler(appUpdateService)

		appUpdates := v1.Group("/app-updates")
		{
			appUpdates.GET("/check", appUpdateHandler.CheckUpdate)
			appUpdates.GET("/latest", appUpdateHandler.GetLatest)
			appUpdates.GET("/history", appUpdateHandler.GetHistory)
			appUpdates.POST("", appUpdateHandler.Upload)
		}
	}

	// 静态文件服务
	router.Static("/uploads", "./uploads")

	return router
}
