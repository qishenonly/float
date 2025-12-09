package response

import "time"

// TransactionResponse 交易响应
type TransactionResponse struct {
	ID              int64                 `json:"id"`
	UserID          int64                 `json:"user_id"`
	Type            string                `json:"type"`
	CategoryID      *int64                `json:"category_id"`
	Category        *CategoryResponse     `json:"category,omitempty"`
	AccountID       *int64                `json:"account_id"`
	Account         *AccountResponse      `json:"account,omitempty"`
	ToAccountID     *int64                `json:"to_account_id"`
	ToAccount       *AccountResponse      `json:"to_account,omitempty"`
	Amount          float64               `json:"amount"`
	Currency        string                `json:"currency"`
	Title           string                `json:"title"`
	Description     string                `json:"description"`
	Location        string                `json:"location"`
	TransactionDate time.Time             `json:"transaction_date"`
	TransactionTime *time.Time            `json:"transaction_time"`
	BillID          *int64                `json:"bill_id"`
	WishlistID      *int64                `json:"wishlist_id"`
	Tags            []string              `json:"tags"`
	Images          []string              `json:"images"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
}

// TransactionListResponse 交易列表响应
type TransactionListResponse struct {
	Total    int64                   `json:"total"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"page_size"`
	Items    []*TransactionResponse  `json:"items"`
}

// TransactionStatisticsResponse 交易统计响应
type TransactionStatisticsResponse struct {
	TotalIncome    float64 `json:"total_income"`
	TotalExpense   float64 `json:"total_expense"`
	NetAmount      float64 `json:"net_amount"`
	TransactionCnt int64   `json:"transaction_cnt"`
}

// MonthlyStatisticsResponse 月度统计响应
type MonthlyStatisticsResponse struct {
	Month          string  `json:"month"`
	TotalIncome    float64 `json:"total_income"`
	TotalExpense   float64 `json:"total_expense"`
	NetAmount      float64 `json:"net_amount"`
	TransactionCnt int64   `json:"transaction_cnt"`
}

// CategoryStatisticsResponse 分类统计响应
type CategoryStatisticsResponse struct {
	CategoryID   int64                 `json:"category_id"`
	Category     *CategoryResponse     `json:"category"`
	TotalAmount  float64               `json:"total_amount"`
	Percentage   float64               `json:"percentage"`
	TransactionCnt int64               `json:"transaction_cnt"`
}

// TransactionDetailsResponse 交易详情响应
type TransactionDetailsResponse struct {
	Transaction    *TransactionResponse           `json:"transaction"`
	RelatedBill    interface{}                    `json:"related_bill,omitempty"`
	RelatedWishlist interface{}                   `json:"related_wishlist,omitempty"`
}

// BulkOperationResponse 批量操作响应
type BulkOperationResponse struct {
	SuccessCount int64  `json:"success_count"`
	FailureCount int64  `json:"failure_count"`
	Errors       []string `json:"errors,omitempty"`
}
