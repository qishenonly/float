package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/dto/response"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/internal/repository"
)

// TransactionService 交易业务逻辑层
type TransactionService interface {
	CreateTransaction(userID int64, req *request.CreateTransactionRequest) (*response.TransactionResponse, error)
	CreateBatchTransactions(userID int64, req *request.BulkCreateTransactionRequest) (*response.BulkOperationResponse, error)
	UpdateTransaction(userID int64, transactionID int64, req *request.UpdateTransactionRequest) (*response.TransactionResponse, error)
	DeleteTransaction(userID int64, transactionID int64) error
	DeleteBatchTransactions(userID int64, req *request.BulkDeleteTransactionRequest) (*response.BulkOperationResponse, error)
	GetTransactionByID(userID int64, transactionID int64) (*response.TransactionResponse, error)
	ListTransactions(userID int64, filters *request.ListTransactionRequest) (*response.TransactionListResponse, error)
	GetTransactionStatistics(userID int64, startDate, endDate time.Time) (*response.TransactionStatisticsResponse, error)
	GetMonthlyStatistics(userID int64, month time.Time) ([]*response.MonthlyStatisticsResponse, error)
	GetCategoryStatistics(userID int64, startDate, endDate time.Time) ([]*response.CategoryStatisticsResponse, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	accountRepo     repository.AccountRepository
	categoryRepo    repository.CategoryRepository
}

// NewTransactionService 创建交易服务实例
func NewTransactionService(
	transactionRepo repository.TransactionRepository,
	accountRepo repository.AccountRepository,
	categoryRepo repository.CategoryRepository,
) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
		categoryRepo:    categoryRepo,
	}
}

// CreateTransaction 创建交易
func (s *transactionService) CreateTransaction(userID int64, req *request.CreateTransactionRequest) (*response.TransactionResponse, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	// 验证账户
	if req.AccountID != nil {
		account, err := s.accountRepo.FindByID(*req.AccountID)
		if err != nil || account.UserID != userID {
			return nil, errors.New("invalid account")
		}
	}

	// 验证转出账户（转账时）
	if req.Type == "transfer" && req.ToAccountID != nil {
		toAccount, err := s.accountRepo.FindByID(*req.ToAccountID)
		if err != nil || toAccount.UserID != userID {
			return nil, errors.New("invalid to_account")
		}
	}

	// 验证分类（系统分类user_id为0，不需要匹配）
	if req.CategoryID != nil {
		category, err := s.categoryRepo.FindByID(*req.CategoryID)
		if err != nil || (category.UserID != userID && category.UserID != 0) {
			return nil, errors.New("invalid category")
		}
	}

	// 设置默认货币
	if req.Currency == "" {
		req.Currency = "CNY"
	}

	// 创建交易对象
	transaction := &models.Transaction{
		UserID:          userID,
		Type:            req.Type,
		CategoryID:      req.CategoryID,
		AccountID:       req.AccountID,
		ToAccountID:     req.ToAccountID,
		Amount:          req.Amount,
		Currency:        req.Currency,
		Title:           req.Title,
		Description:     req.Description,
		Location:        req.Location,
		TransactionDate: req.TransactionDate,
		TransactionTime: req.TransactionTime,
		BillID:          req.BillID,
		WishlistID:      req.WishlistID,
		Tags:            models.JSONArray(req.Tags),
		Images:          models.JSONArray(req.Images),
	}

	// 保存交易并更新账户余额
	if err := s.transactionRepo.Create(transaction); err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	// 更新账户余额
	if transaction.Type == "expense" || transaction.Type == "income" {
		if req.AccountID != nil {
			account, _ := s.accountRepo.FindByID(*req.AccountID)
			if transaction.Type == "income" {
				account.Balance += transaction.Amount
			} else {
				account.Balance -= transaction.Amount
			}
			if err := s.accountRepo.Update(account); err != nil {
				return nil, fmt.Errorf("failed to update account: %w", err)
			}
		}
	} else if transaction.Type == "transfer" {
		// 处理转账：从源账户扣除，到目标账户增加
		if req.AccountID != nil {
			fromAccount, _ := s.accountRepo.FindByID(*req.AccountID)
			fromAccount.Balance -= transaction.Amount
			if err := s.accountRepo.Update(fromAccount); err != nil {
				return nil, fmt.Errorf("failed to update account: %w", err)
			}
		}

		if req.ToAccountID != nil {
			toAccount, _ := s.accountRepo.FindByID(*req.ToAccountID)
			toAccount.Balance += transaction.Amount
			if err := s.accountRepo.Update(toAccount); err != nil {
				return nil, fmt.Errorf("failed to update account: %w", err)
			}
		}
	}

	// 获取完整的交易信息
	fullTransaction, _ := s.transactionRepo.FindByID(transaction.ID)
	return s.toTransactionResponse(fullTransaction), nil
}

// CreateBatchTransactions 批量创建交易
func (s *transactionService) CreateBatchTransactions(userID int64, req *request.BulkCreateTransactionRequest) (*response.BulkOperationResponse, error) {
	if req == nil || len(req.Transactions) == 0 {
		return nil, errors.New("request cannot be empty")
	}

	result := &response.BulkOperationResponse{
		Errors: []string{},
	}

	// 逐条创建并记录错误
	for _, transReq := range req.Transactions {
		_, err := s.CreateTransaction(userID, &transReq)
		if err != nil {
			result.FailureCount++
			result.Errors = append(result.Errors, err.Error())
		} else {
			result.SuccessCount++
		}
	}

	return result, nil
}

// UpdateTransaction 更新交易
func (s *transactionService) UpdateTransaction(userID int64, transactionID int64, req *request.UpdateTransactionRequest) (*response.TransactionResponse, error) {
	// 获取原交易
	oldTransaction, err := s.transactionRepo.FindByID(transactionID)
	if err != nil || oldTransaction.UserID != userID {
		return nil, errors.New("transaction not found")
	}

	// 更新字段
	if req.Type != "" {
		oldTransaction.Type = req.Type
	}
	if req.CategoryID != nil {
		oldTransaction.CategoryID = req.CategoryID
	}
	if req.AccountID != nil {
		oldTransaction.AccountID = req.AccountID
	}
	if req.ToAccountID != nil {
		oldTransaction.ToAccountID = req.ToAccountID
	}
	if req.Amount != 0 {
		oldTransaction.Amount = req.Amount
	}
	if req.Currency != "" {
		oldTransaction.Currency = req.Currency
	}
	if req.Title != "" {
		oldTransaction.Title = req.Title
	}
	if req.Description != "" {
		oldTransaction.Description = req.Description
	}
	if req.Location != "" {
		oldTransaction.Location = req.Location
	}
	if !req.TransactionDate.IsZero() {
		oldTransaction.TransactionDate = req.TransactionDate
	}
	if req.TransactionTime != nil {
		oldTransaction.TransactionTime = req.TransactionTime
	}
	if len(req.Tags) > 0 {
		oldTransaction.Tags = models.JSONArray(req.Tags)
	}
	if len(req.Images) > 0 {
		oldTransaction.Images = models.JSONArray(req.Images)
	}

	if err := s.transactionRepo.Update(oldTransaction); err != nil {
		return nil, fmt.Errorf("failed to update transaction: %w", err)
	}

	fullTransaction, _ := s.transactionRepo.FindByID(transactionID)
	return s.toTransactionResponse(fullTransaction), nil
}

// DeleteTransaction 删除交易
func (s *transactionService) DeleteTransaction(userID int64, transactionID int64) error {
	transaction, err := s.transactionRepo.FindByID(transactionID)
	if err != nil || transaction.UserID != userID {
		return errors.New("transaction not found")
	}

	// 恢复账户余额
	if transaction.Type == "expense" || transaction.Type == "income" {
		if transaction.AccountID != nil {
			account, _ := s.accountRepo.FindByID(*transaction.AccountID)
			if transaction.Type == "income" {
				account.Balance -= transaction.Amount
			} else {
				account.Balance += transaction.Amount
			}
			_ = s.accountRepo.Update(account)
		}
	} else if transaction.Type == "transfer" {
		// 恢复转账的两个账户
		if transaction.AccountID != nil {
			fromAccount, _ := s.accountRepo.FindByID(*transaction.AccountID)
			fromAccount.Balance += transaction.Amount
			_ = s.accountRepo.Update(fromAccount)
		}

		if transaction.ToAccountID != nil {
			toAccount, _ := s.accountRepo.FindByID(*transaction.ToAccountID)
			toAccount.Balance -= transaction.Amount
			_ = s.accountRepo.Update(toAccount)
		}
	}

	return s.transactionRepo.Delete(transactionID)
}

// DeleteBatchTransactions 批量删除交易
func (s *transactionService) DeleteBatchTransactions(userID int64, req *request.BulkDeleteTransactionRequest) (*response.BulkOperationResponse, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, errors.New("request cannot be empty")
	}

	result := &response.BulkOperationResponse{
		Errors: []string{},
	}

	for _, id := range req.IDs {
		err := s.DeleteTransaction(userID, id)
		if err != nil {
			result.FailureCount++
			result.Errors = append(result.Errors, err.Error())
		} else {
			result.SuccessCount++
		}
	}

	return result, nil
}

// GetTransactionByID 获取交易详情
func (s *transactionService) GetTransactionByID(userID int64, transactionID int64) (*response.TransactionResponse, error) {
	transaction, err := s.transactionRepo.FindByID(transactionID)
	if err != nil || transaction.UserID != userID {
		return nil, errors.New("transaction not found")
	}

	return s.toTransactionResponse(transaction), nil
}

// ListTransactions 查询交易列表
func (s *transactionService) ListTransactions(userID int64, filters *request.ListTransactionRequest) (*response.TransactionListResponse, error) {
	if filters == nil {
		filters = &request.ListTransactionRequest{
			Page:     1,
			PageSize: 20,
		}
	}

	transactions, total, err := s.transactionRepo.FindByUserID(userID, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to list transactions: %w", err)
	}

	items := make([]*response.TransactionResponse, len(transactions))
	for i, t := range transactions {
		items[i] = s.toTransactionResponse(t)
	}

	return &response.TransactionListResponse{
		Total:    total,
		Page:     filters.Page,
		PageSize: filters.PageSize,
		Items:    items,
	}, nil
}

// GetTransactionStatistics 获取交易统计
func (s *transactionService) GetTransactionStatistics(userID int64, startDate, endDate time.Time) (*response.TransactionStatisticsResponse, error) {
	// 这里需要通过SQL查询进行统计，实际实现需要在repository中添加专门的统计方法
	// 这里给出简化版本

	filters := &request.ListTransactionRequest{
		StartDate: startDate,
		EndDate:   endDate,
		Page:      1,
		PageSize:  1000,
	}

	transactions, _, err := s.transactionRepo.FindByUserID(userID, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to get statistics: %w", err)
	}

	stats := &response.TransactionStatisticsResponse{}
	for _, t := range transactions {
		stats.TransactionCnt++
		if t.Type == "income" {
			stats.TotalIncome += t.Amount
		} else if t.Type == "expense" {
			stats.TotalExpense += t.Amount
		}
	}
	stats.NetAmount = stats.TotalIncome - stats.TotalExpense

	return stats, nil
}

// GetMonthlyStatistics 获取月度统计
func (s *transactionService) GetMonthlyStatistics(userID int64, month time.Time) ([]*response.MonthlyStatisticsResponse, error) {
	// 实现月度统计逻辑
	startDate := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
	endDate := startDate.AddDate(0, 1, -1)

	stats, err := s.GetTransactionStatistics(userID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return []*response.MonthlyStatisticsResponse{
		{
			Month:          month.Format("2006-01"),
			TotalIncome:    stats.TotalIncome,
			TotalExpense:   stats.TotalExpense,
			NetAmount:      stats.NetAmount,
			TransactionCnt: stats.TransactionCnt,
		},
	}, nil
}

// GetCategoryStatistics 获取分类统计
func (s *transactionService) GetCategoryStatistics(userID int64, startDate, endDate time.Time) ([]*response.CategoryStatisticsResponse, error) {
	// 实现分类统计逻辑
	filters := &request.ListTransactionRequest{
		Type:      "expense",
		StartDate: startDate,
		EndDate:   endDate,
		Page:      1,
		PageSize:  1000,
	}

	transactions, _, err := s.transactionRepo.FindByUserID(userID, filters)
	if err != nil {
		return nil, err
	}

	// 按分类汇总
	categoryMap := make(map[int64]*response.CategoryStatisticsResponse)
	var totalExpense float64

	for _, t := range transactions {
		totalExpense += t.Amount
		if t.CategoryID == nil {
			continue
		}

		catID := *t.CategoryID
		if _, exists := categoryMap[catID]; !exists {
			categoryMap[catID] = &response.CategoryStatisticsResponse{
				CategoryID: catID,
				Category:   s.toCategoryResponse(t.Category),
			}
		}

		categoryMap[catID].TotalAmount += t.Amount
		categoryMap[catID].TransactionCnt++
	}

	// 计算百分比
	result := make([]*response.CategoryStatisticsResponse, 0)
	for _, stat := range categoryMap {
		if totalExpense > 0 {
			stat.Percentage = stat.TotalAmount / totalExpense * 100
		}
		result = append(result, stat)
	}

	return result, nil
}

// 辅助方法：将Transaction模型转换为Response
func (s *transactionService) toTransactionResponse(t *models.Transaction) *response.TransactionResponse {
	if t == nil {
		return nil
	}

	resp := &response.TransactionResponse{
		ID:              t.ID,
		UserID:          t.UserID,
		Type:            t.Type,
		CategoryID:      t.CategoryID,
		AccountID:       t.AccountID,
		ToAccountID:     t.ToAccountID,
		Amount:          t.Amount,
		Currency:        t.Currency,
		Title:           t.Title,
		Description:     t.Description,
		Location:        t.Location,
		TransactionDate: t.TransactionDate,
		TransactionTime: t.TransactionTime,
		BillID:          t.BillID,
		WishlistID:      t.WishlistID,
		Tags:            t.Tags,
		Images:          t.Images,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
	}

	if t.Category != nil {
		resp.Category = s.toCategoryResponse(t.Category)
	}

	if t.Account != nil {
		resp.Account = s.toAccountResponse(t.Account)
	}

	if t.ToAccount != nil {
		resp.ToAccount = s.toAccountResponse(t.ToAccount)
	}

	return resp
}

// 辅助方法：将Category模型转换为Response
func (s *transactionService) toCategoryResponse(c *models.Category) *response.CategoryResponse {
	if c == nil {
		return nil
	}

	return &response.CategoryResponse{
		ID:           c.ID,
		UserID:       c.UserID,
		Type:         c.Type,
		Name:         c.Name,
		Icon:         c.Icon,
		Color:        c.Color,
		DisplayOrder: c.DisplayOrder,
		IsSystem:     c.IsSystem,
		IsActive:     c.IsActive,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

// 辅助方法：将Account模型转换为Response
func (s *transactionService) toAccountResponse(a *models.Account) *response.AccountResponse {
	if a == nil {
		return nil
	}

	return &response.AccountResponse{
		ID:             a.ID,
		UserID:         a.UserID,
		AccountType:    a.AccountType,
		AccountName:    a.AccountName,
		AccountNumber:  a.AccountNumber,
		Icon:           a.Icon,
		Color:          a.Color,
		Balance:        a.Balance,
		InitialBalance: a.InitialBalance,
		IncludeInTotal: a.IncludeInTotal,
		DisplayOrder:   a.DisplayOrder,
		IsActive:       a.IsActive,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
	}
}
