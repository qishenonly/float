package request

import "time"

// CreateTransactionRequest 创建交易请求
type CreateTransactionRequest struct {
	Type            string     `json:"type" binding:"required,oneof=expense income transfer"`
	CategoryID      *int64     `json:"category_id"`
	AccountID       *int64     `json:"account_id" binding:"required_if=Type expense,required_if=Type income"`
	ToAccountID     *int64     `json:"to_account_id" binding:"required_if=Type transfer"`
	Amount          float64    `json:"amount" binding:"required,gt=0"`
	Currency        string     `json:"currency" binding:"omitempty,len=3"`
	Title           string     `json:"title" binding:"max=200"`
	Description     string     `json:"description"`
	Location        string     `json:"location" binding:"max=200"`
	TransactionDate time.Time  `json:"transaction_date" binding:"required"`
	TransactionTime *time.Time `json:"transaction_time"`
	BillID          *int64     `json:"bill_id"`
	WishlistID      *int64     `json:"wishlist_id"`
	Tags            []string   `json:"tags"`
	Images          []string   `json:"images"`
}

// UpdateTransactionRequest 更新交易请求
type UpdateTransactionRequest struct {
	Type            string     `json:"type" binding:"omitempty,oneof=expense income transfer"`
	CategoryID      *int64     `json:"category_id"`
	AccountID       *int64     `json:"account_id"`
	ToAccountID     *int64     `json:"to_account_id"`
	Amount          float64    `json:"amount" binding:"omitempty,gt=0"`
	Currency        string     `json:"currency" binding:"omitempty,len=3"`
	Title           string     `json:"title" binding:"omitempty,max=200"`
	Description     string     `json:"description"`
	Location        string     `json:"location" binding:"omitempty,max=200"`
	TransactionDate time.Time  `json:"transaction_date"`
	TransactionTime *time.Time `json:"transaction_time"`
	BillID          *int64     `json:"bill_id"`
	WishlistID      *int64     `json:"wishlist_id"`
	Tags            []string   `json:"tags"`
	Images          []string   `json:"images"`
}

// ListTransactionRequest 查询交易列表请求
type ListTransactionRequest struct {
	Type           string    `form:"type" binding:"omitempty,oneof=expense income transfer"`
	CategoryID     *int64    `form:"category_id"`
	AccountID      *int64    `form:"account_id"`
	StartDate      time.Time `form:"start_date" binding:"omitempty"`
	EndDate        time.Time `form:"end_date" binding:"omitempty"`
	SearchKeyword  string    `form:"search_keyword" binding:"max=100"`
	SortBy         string    `form:"sort_by" binding:"omitempty,oneof=date amount"`
	SortOrder      string    `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	Page           int       `form:"page" binding:"omitempty,min=1"`
	PageSize       int       `form:"page_size" binding:"omitempty,min=1,max=100"`
}

// BulkCreateTransactionRequest 批量创建交易请求
type BulkCreateTransactionRequest struct {
	Transactions []CreateTransactionRequest `json:"transactions" binding:"required,min=1,max=100"`
}

// BulkDeleteTransactionRequest 批量删除交易请求
type BulkDeleteTransactionRequest struct {
	IDs []int64 `json:"ids" binding:"required,min=1,max=1000"`
}
