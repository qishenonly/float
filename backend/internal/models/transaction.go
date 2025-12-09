package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Transaction 交易记录表
type Transaction struct {
	ID                int64       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID            int64       `gorm:"not null;index:idx_user_date;index:idx_user_type" json:"user_id"`
	Type              string      `gorm:"type:enum('expense','income','transfer');not null;index:idx_user_type" json:"type"`
	CategoryID        *int64      `gorm:"index:idx_category" json:"category_id"`
	AccountID         *int64      `gorm:"index:idx_account" json:"account_id"`
	ToAccountID       *int64      `json:"to_account_id"`
	Amount            float64     `gorm:"type:decimal(15,2);not null" json:"amount"`
	Currency          string      `gorm:"size:10;default:'CNY'" json:"currency"`
	Title             string      `gorm:"size:200" json:"title"`
	Description       string      `gorm:"type:text" json:"description"`
	Location          string      `gorm:"size:200" json:"location"`
	TransactionDate   time.Time   `gorm:"type:date;not null;index:idx_date;index:idx_user_date" json:"transaction_date"`
	TransactionTime   *time.Time  `json:"transaction_time"`
	BillID            *int64      `json:"bill_id"`
	WishlistID        *int64      `json:"wishlist_id"`
	Tags              JSONArray   `gorm:"type:json" json:"tags"`
	Images            JSONArray   `gorm:"type:json" json:"images"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`

	// 关联关系
	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Account  *Account  `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	ToAccount *Account `gorm:"foreignKey:ToAccountID" json:"to_account,omitempty"`
}

// TableName 表名
func (Transaction) TableName() string {
	return "transactions"
}

// JSONArray 用于处理JSON数组字段
type JSONArray []string

// Value 实现driver.Valuer接口
func (ja JSONArray) Value() (driver.Value, error) {
	return json.Marshal(ja)
}

// Scan 实现sql.Scanner接口
func (ja *JSONArray) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}

	return json.Unmarshal(bytes, &ja)
}

// TransactionType 交易类型常量
type TransactionType string

const (
	TransactionTypeExpense  TransactionType = "expense"
	TransactionTypeIncome   TransactionType = "income"
	TransactionTypeTransfer TransactionType = "transfer"
)
