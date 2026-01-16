package request

// CreateAccountRequest 创建账户请求
type CreateAccountRequest struct {
	AccountType    string  `json:"account_type" binding:"required,oneof=bank alipay wechat cash credit other"`
	AccountName    string  `json:"account_name" binding:"required,max=100"`
	AccountNumber  string  `json:"account_number" binding:"max=50"`
	Icon           string  `json:"icon" binding:"max=50"`
	Color          string  `json:"color" binding:"max=20"`
	InitialBalance float64 `json:"initial_balance"`
	IncludeInTotal *bool   `json:"include_in_total"` // 使用指针以支持false值
	DisplayOrder   int     `json:"display_order"`
}

// UpdateAccountRequest 更新账户请求
type UpdateAccountRequest struct {
	AccountName    *string  `json:"account_name" binding:"omitempty,max=100"`
	AccountNumber  *string  `json:"account_number" binding:"omitempty,max=50"`
	Icon           *string  `json:"icon" binding:"omitempty,max=50"`
	Color          *string  `json:"color" binding:"omitempty,max=20"`
	InitialBalance *float64 `json:"initial_balance"`
	IncludeInTotal *bool    `json:"include_in_total"`
	DisplayOrder   *int     `json:"display_order"`
	IsActive       *bool    `json:"is_active"`
	Balance        *float64 `json:"balance"` // 允许直接修改余额（修正用）
}
