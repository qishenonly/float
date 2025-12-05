package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/api/routes"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/pkg/cache"
	"github.com/qiuhaonan/float-backend/pkg/database"
	"github.com/qiuhaonan/float-backend/pkg/logger"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置
	if err := initConfig(); err != nil {
		log.Fatal("Failed to initialize config:", err)
	}

	// 初始化日志
	logger.Init()

	// 初始化数据库
	if err := database.Init(); err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}
	defer database.Close()

	// 自动迁移数据库表
	logger.Info("Running database migrations...")
	if err := database.AutoMigrate(&models.User{}); err != nil {
		logger.Fatal("Failed to migrate database:", err)
	}
	logger.Info("Database migrations completed")

	// 初始化 Redis
	if err := cache.Init(); err != nil {
		logger.Fatal("Failed to initialize Redis:", err)
	}
	defer cache.Close()

	// 设置 Gin 模式
	if viper.GetString("app.mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化路由
	router := routes.SetupRouter()

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")),
		Handler:        router,
		ReadTimeout:    time.Duration(viper.GetInt("server.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("server.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	// 启动服务器（goroutine）
	go func() {
		logger.Infof("Server starting on %s:%s", viper.GetString("server.host"), viper.GetString("server.port"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server:", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", err)
	}

	logger.Info("Server exited")
}

func initConfig() error {
	viper.SetConfigName("config.dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	// 设置默认值
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.read_timeout", 10)
	viper.SetDefault("server.write_timeout", 10)
	viper.SetDefault("app.mode", "debug")

	// 读取环境变量
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn("Config file not found, using defaults")
			return nil
		}
		return err
	}

	return nil
}
