package models

import (
	"time"
)

// Account 资金账户表
type Account struct {
	ID             int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         int64     `gorm:"not null;index:idx_user_id" json:"user_id"`
	AccountType    string    `gorm:"type:enum('bank','alipay','wechat','cash','other');not null" json:"account_type"`
	AccountName    string    `gorm:"size:100;not null" json:"account_name"`
	AccountNumber  string    `gorm:"size:50" json:"account_number"`
	Icon           string    `gorm:"size:50" json:"icon"`
	Color          string    `gorm:"size:20" json:"color"`
	Balance        float64   `gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	InitialBalance float64   `gorm:"type:decimal(15,2);default:0.00" json:"initial_balance"`
	IncludeInTotal bool      `gorm:"default:true" json:"include_in_total"`
	DisplayOrder   int       `gorm:"default:0" json:"display_order"`
	IsActive       bool      `gorm:"default:true;index:idx_active" json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// TableName 表名
func (Account) TableName() string {
	return "accounts"
}
