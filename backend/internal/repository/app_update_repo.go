package repository

import (
	"github.com/qiuhaonan/float-backend/internal/models"
	"gorm.io/gorm"
)

type AppUpdateRepository struct {
	db *gorm.DB
}

func NewAppUpdateRepository(db *gorm.DB) *AppUpdateRepository {
	return &AppUpdateRepository{db: db}
}

// GetLatestVersion 获取指定平台最新发布的版本
func (r *AppUpdateRepository) GetLatestVersion(platform string) (*models.AppUpdate, error) {
	var update models.AppUpdate
	err := r.db.Where("platform = ? AND status = 'released'", platform).
		Order("version_code DESC").
		First(&update).Error
	if err != nil {
		return nil, err
	}
	return &update, nil
}

// GetUpdateHistory 获取更新历史
func (r *AppUpdateRepository) GetUpdateHistory(platform string, limit int) ([]models.AppUpdate, error) {
	var updates []models.AppUpdate
	err := r.db.Where("platform = ? AND status = 'released'", platform).
		Order("version_code DESC").
		Limit(limit).
		Find(&updates).Error
	return updates, err
}

// Create 创建更新记录
func (r *AppUpdateRepository) Create(update *models.AppUpdate) error {
	return r.db.Create(update).Error
}

// GetByID 获取指定ID的更新详情
func (r *AppUpdateRepository) GetByID(id int64) (*models.AppUpdate, error) {
	var update models.AppUpdate
	err := r.db.First(&update, id).Error
	if err != nil {
		return nil, err
	}
	return &update, nil
}
