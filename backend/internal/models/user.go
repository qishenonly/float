package models

import (
	"time"
)

// User 用户表
type User struct {
	ID              int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username        string     `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Email           string     `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Phone           *string    `gorm:"size:20;uniqueIndex" json:"phone,omitempty"`
	PasswordHash    string     `gorm:"size:255;not null" json:"-"`
	AvatarURL       string     `gorm:"size:500" json:"avatar_url,omitempty"`
	DisplayName     string     `gorm:"size:50" json:"display_name,omitempty"`
	Verified        bool       `gorm:"default:false" json:"verified"`
	Currency        string     `gorm:"size:10;default:'CNY'" json:"currency"`
	Theme           string     `gorm:"size:20;default:'light'" json:"theme"`
	Language        string     `gorm:"size:10;default:'zh-CN'" json:"language"`
	DarkMode        bool       `gorm:"default:false" json:"dark_mode"`
	GestureLock     bool       `gorm:"default:true" json:"gesture_lock"`
	ContinuousDays  int        `gorm:"default:0" json:"continuous_days"`
	TotalRecords    int        `gorm:"default:0" json:"total_records"`
	TotalBadges     int        `gorm:"default:0" json:"total_badges"`
	MembershipLevel string     `gorm:"size:20;default:'FREE'" json:"membership_level"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	LastLoginAt     *time.Time `json:"last_login_at,omitempty"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}
