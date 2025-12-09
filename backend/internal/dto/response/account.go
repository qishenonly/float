package response

import "time"

// AccountResponse 账户响应
type AccountResponse struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"user_id"`
	AccountType    string    `json:"account_type"`
	AccountName    string    `json:"account_name"`
	AccountNumber  string    `json:"account_number"`
	Icon           string    `json:"icon"`
	Color          string    `json:"color"`
	Balance        float64   `json:"balance"`
	InitialBalance float64   `json:"initial_balance"`
	IncludeInTotal bool      `json:"include_in_total"`
	DisplayOrder   int       `json:"display_order"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// AccountBalanceResponse 账户余额汇总响应
type AccountBalanceResponse struct {
	TotalBalance float64 `json:"total_balance"`
	AssetBalance float64 `json:"asset_balance"` // 净资产（不含负债）
	DebtBalance  float64 `json:"debt_balance"`  // 负债总额
}
