package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/internal/service"
	"github.com/qiuhaonan/float-backend/pkg/database"
	"github.com/qiuhaonan/float-backend/pkg/logger"
)

type AppUpdateHandler struct {
	service *service.AppUpdateService
}

func NewAppUpdateHandler(service *service.AppUpdateService) *AppUpdateHandler {
	return &AppUpdateHandler{service: service}
}

// CheckUpdate 检查更新
// @Summary 检查应用更新
// @Description 检查指定平台和版本是否有新版本
// @Tags AppUpdates
// @Accept json
// @Produce json
// @Param platform query string true "平台 (android, ios, web)"
// @Param version_code query int true "当前版本代码"
// @Success 200 {object} response.CheckUpdateResponse
// @Router /api/v1/app-updates/check [get]
func (h *AppUpdateHandler) CheckUpdate(c *gin.Context) {
	platform := c.Query("platform")
	versionCodeStr := c.Query("version_code")

	if platform == "" || versionCodeStr == "" {
		logger.Error("[Handler][检查更新] 缺少必要参数 | platform: " + platform + " | version_code: " + versionCodeStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing platform or version_code"})
		return
	}

	versionCode, err := strconv.Atoi(versionCodeStr)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][检查更新] 版本号格式错误 | platform: %s | version_code: %s", platform, versionCodeStr))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid version_code"})
		return
	}

	logger.Info(fmt.Sprintf("[Handler][检查更新] 检查更新请求 | 平台: %s | 当前版本: %d", platform, versionCode))

	resp, err := h.service.CheckUpdate(platform, versionCode)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][检查更新] 检查失败 | 平台: %s | 版本: %d | 错误: %v", platform, versionCode, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	logger.Info(fmt.Sprintf("[Handler][检查更新] 检查成功 | 平台: %s | 当前版本: %d | 是否需要更新: %v", platform, versionCode, resp))
	c.JSON(http.StatusOK, resp)
}

// GetLatest 获取最新版本
// @Summary 获取最新版本信息
// @Description 获取指定平台的最新版本详情
// @Tags AppUpdates
// @Accept json
// @Produce json
// @Param platform query string true "平台 (android, ios, web)"
// @Success 200 {object} response.AppUpdateResponse
// @Router /api/v1/app-updates/latest [get]
func (h *AppUpdateHandler) GetLatest(c *gin.Context) {
	platform := c.Query("platform")
	if platform == "" {
		logger.Error("[Handler][获取最新版本] 缺少platform参数")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing platform"})
		return
	}

	logger.Info(fmt.Sprintf("[Handler][获取最新版本] 获取最新版本请求 | 平台: %s", platform))

	resp, err := h.service.GetLatest(platform)
	if err != nil {
		if err == database.ErrRecordNotFound { // Assuming gorm.ErrRecordNotFound is handled or wrapped
			logger.Warn(fmt.Sprintf("[Handler][获取最新版本] 未找到版本 | 平台: %s", platform))
			c.JSON(http.StatusNotFound, gin.H{"error": "No version found"})
			return
		}
		logger.Error(fmt.Sprintf("[Handler][获取最新版本] 获取失败 | 平台: %s | 错误: %v", platform, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	logger.Info(fmt.Sprintf("[Handler][获取最新版本] 获取成功 | 平台: %s | 版本代码: %v", platform, resp))
	c.JSON(http.StatusOK, resp)
}

// GetHistory 获取更新历史
// @Summary 获取更新历史
// @Description 获取指定平台的更新历史记录
// @Tags AppUpdates
// @Accept json
// @Produce json
// @Param platform query string true "平台 (android, ios, web)"
// @Success 200 {array} response.AppUpdateResponse
// @Router /api/v1/app-updates/history [get]
func (h *AppUpdateHandler) GetHistory(c *gin.Context) {
	platform := c.Query("platform")
	if platform == "" {
		logger.Error("[Handler][获取更新历史] 缺少platform参数")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing platform"})
		return
	}

	logger.Info(fmt.Sprintf("[Handler][获取更新历史] 获取更新历史请求 | 平台: %s", platform))

	resp, err := h.service.GetHistory(platform)
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][获取更新历史] 获取失败 | 平台: %s | 错误: %v", platform, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	logger.Info(fmt.Sprintf("[Handler][获取更新历史] 获取成功 | 平台: %s | 版本数: %v", platform, len(resp)))
	c.JSON(http.StatusOK, resp)
}

// Upload 上传更新包
// @Summary 上传应用更新包
// @Description 上传APK文件并创建更新记录
// @Tags AppUpdates
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "APK文件"
// @Param version_code formData int true "版本代码"
// @Param version_name formData string true "版本名称"
// @Param platform formData string true "平台 (android)"
// @Param update_type formData string true "更新类型 (major, minor, patch)"
// @Param title formData string true "更新标题"
// @Param description formData string true "更新描述"
// @Param changelog formData string false "变更日志 (JSON)"
// @Param is_force_update formData bool false "是否强制更新"
// @Success 200 {object} response.AppUpdateResponse
// @Router /api/v1/app-updates [post]
func (h *AppUpdateHandler) Upload(c *gin.Context) {
	// 1. 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error(fmt.Sprintf("[Handler][上传更新] 文件获取失败 | 错误: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// 2. 获取其他参数
	versionCodeStr := c.PostForm("version_code")
	versionName := c.PostForm("version_name")
	platform := c.PostForm("platform")
	updateType := c.PostForm("update_type")
	title := c.PostForm("title")
	description := c.PostForm("description")
	changelogStr := c.PostForm("changelog")
	isForceUpdateStr := c.PostForm("is_force_update")

	if versionCodeStr == "" || versionName == "" || platform == "" || title == "" {
		logger.Error(fmt.Sprintf("[Handler][上传更新] 缺少必要参数 | version_code: %s | version_name: %s | platform: %s | title: %s", versionCodeStr, versionName, platform, title))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	logger.Info(fmt.Sprintf("[Handler][上传更新] 上传更新包请求 | 平台: %s | 版本: %s | 文件大小: %d", platform, versionName, file.Size))

	versionCode, _ := strconv.Atoi(versionCodeStr)
	isForceUpdate, _ := strconv.ParseBool(isForceUpdateStr)

	// 3. 保存文件
	// 确保目录存在: uploads/apk/<platform>/
	uploadDir := fmt.Sprintf("uploads/apk/%s", platform)
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		logger.Error("Failed to create upload directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// 文件名: <version_code>_<filename>
	filename := fmt.Sprintf("%d_%s", versionCode, file.Filename)
	filepath := fmt.Sprintf("%s/%s", uploadDir, filename)

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		logger.Error("Failed to save file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// 4. 创建记录
	// 构造相对路径 URL
	downloadURL := fmt.Sprintf("/%s", filepath)

	appUpdate := &models.AppUpdate{
		VersionCode:   versionCode,
		VersionName:   versionName,
		Platform:      platform,
		UpdateType:    updateType,
		IsForceUpdate: isForceUpdate,
		Title:         title,
		Description:   description,
		Changelog:     json.RawMessage(changelogStr),
		DownloadURL:   downloadURL,
		FileSize:      file.Size,
		Status:        "released", // 默认直接发布
	}

	if err := h.service.CreateAppUpdate(appUpdate); err != nil {
		logger.Error(fmt.Sprintf("[Handler][上传更新] 创建记录失败 | 平台: %s | 版本: %s | 错误: %v", platform, versionName, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create record"})
		return
	}

	logger.Info(fmt.Sprintf("[Handler][上传更新] 上传成功 | 平台: %s | 版本: %s | 文件路径: %s", platform, versionName, filepath))
	c.JSON(http.StatusOK, gin.H{
		"message": "Upload successful",
		"data":    appUpdate,
	})
}
