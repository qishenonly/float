package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// AppUpdate 软件更新模型
type AppUpdate struct {
	ID                  int64           `gorm:"primaryKey;autoIncrement" json:"id"`
	VersionCode         int             `gorm:"not null;index:idx_version" json:"version_code"`
	VersionName         string          `gorm:"size:20;not null" json:"version_name"`
	Platform            string          `gorm:"type:enum('android','ios','web','all');not null;index:idx_platform_status" json:"platform"`
	UpdateType          string          `gorm:"type:enum('major','minor','patch','hotfix');not null" json:"update_type"`
	IsForceUpdate       bool            `gorm:"default:false" json:"is_force_update"`
	MinSupportedVersion string          `gorm:"size:20" json:"min_supported_version"`
	Title               string          `gorm:"size:200;not null" json:"title"`
	Description         string          `gorm:"type:text;not null" json:"description"`
	Changelog           json.RawMessage `gorm:"type:json" json:"changelog"`
	DownloadURL         string          `gorm:"size:500" json:"download_url"`
	FileSize            int64           `json:"file_size"`
	FileHash            string          `gorm:"size:64" json:"file_hash"`
	ReleaseNotesURL     string          `gorm:"size:500" json:"release_notes_url"`
	Status              string          `gorm:"type:enum('draft','beta','released','deprecated');default:'released';index:idx_platform_status" json:"status"`
	ReleaseDate         *time.Time      `gorm:"index:idx_release_date" json:"release_date"`
	CreatedAt           time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (AppUpdate) TableName() string {
	return "app_updates"
}

// BeforeCreate 创建前钩子
func (u *AppUpdate) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ReleaseDate == nil {
		now := time.Now()
		u.ReleaseDate = &now
	}
	return
}
