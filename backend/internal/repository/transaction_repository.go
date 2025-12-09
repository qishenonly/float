package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/pkg/database"
	"gorm.io/gorm"
)

// TransactionRepository 交易仓库接口
type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	CreateBatch(transactions []*models.Transaction) error
	FindByID(id int64) (*models.Transaction, error)
	FindByUserID(userID int64, filters *request.ListTransactionRequest) ([]*models.Transaction, int64, error)
	Update(transaction *models.Transaction) error
	Delete(id int64) error
	DeleteBatch(ids []int64) error
	GetTotalBalance(userID int64, filters *request.ListTransactionRequest) (*models.Transaction, error)
	GetMonthlyStatistics(userID int64, month time.Time) (*models.Transaction, error)
	GetCategoryStatistics(userID int64, startDate, endDate time.Time) ([]*models.Transaction, error)
	GetDateRangeStatistics(userID int64, startDate, endDate time.Time) (*models.Transaction, error)
}

type transactionRepository struct{}

// NewTransactionRepository 创建交易仓库实例
func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

// Create 创建单条交易
func (r *transactionRepository) Create(transaction *models.Transaction) error {
	if err := database.DB.Create(transaction).Error; err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}
	return nil
}

// CreateBatch 批量创建交易
func (r *transactionRepository) CreateBatch(transactions []*models.Transaction) error {
	if len(transactions) == 0 {
		return errors.New("transactions list cannot be empty")
	}

	if err := database.DB.CreateInBatches(transactions, 100).Error; err != nil {
		return fmt.Errorf("failed to create transactions in batch: %w", err)
	}
	return nil
}

// FindByID 根据ID查找交易
func (r *transactionRepository) FindByID(id int64) (*models.Transaction, error) {
	var transaction models.Transaction
	err := database.DB.
		Preload("Category").
		Preload("Account").
		Preload("ToAccount").
		Where("id = ?", id).
		First(&transaction).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		return nil, fmt.Errorf("failed to find transaction: %w", err)
	}

	return &transaction, nil
}

// FindByUserID 根据用户ID查找交易列表（带分页和过滤）
func (r *transactionRepository) FindByUserID(userID int64, filters *request.ListTransactionRequest) ([]*models.Transaction, int64, error) {
	query := database.DB.Where("user_id = ?", userID)

	// 应用过滤条件
	if filters.Type != "" {
		query = query.Where("type = ?", filters.Type)
	}

	if filters.CategoryID != nil {
		query = query.Where("category_id = ?", *filters.CategoryID)
	}

	if filters.AccountID != nil {
		query = query.Where("account_id = ?", *filters.AccountID)
	}

	if !filters.StartDate.IsZero() {
		query = query.Where("DATE(transaction_date) >= ?", filters.StartDate.Format("2006-01-02"))
	}

	if !filters.EndDate.IsZero() {
		query = query.Where("DATE(transaction_date) <= ?", filters.EndDate.Format("2006-01-02"))
	}

	if filters.SearchKeyword != "" {
		query = query.Where("title LIKE ? OR description LIKE ? OR location LIKE ?",
			"%"+filters.SearchKeyword+"%",
			"%"+filters.SearchKeyword+"%",
			"%"+filters.SearchKeyword+"%")
	}

	// 设置默认分页
	if filters.Page == 0 {
		filters.Page = 1
	}
	if filters.PageSize == 0 {
		filters.PageSize = 20
	}
	if filters.PageSize > 100 {
		filters.PageSize = 100
	}

	// 获取总数
	var total int64
	if err := query.Model(&models.Transaction{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count transactions: %w", err)
	}

	// 排序
	sortBy := "transaction_date"
	if filters.SortBy == "amount" {
		sortBy = "amount"
	}
	sortOrder := "DESC"
	if filters.SortOrder == "asc" {
		sortOrder = "ASC"
	}

	// 查询数据
	var transactions []*models.Transaction
	offset := (filters.Page - 1) * filters.PageSize

	err := query.
		Preload("Category").
		Preload("Account").
		Preload("ToAccount").
		Order(fmt.Sprintf("%s %s", sortBy, sortOrder)).
		Offset(offset).
		Limit(filters.PageSize).
		Find(&transactions).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find transactions: %w", err)
	}

	return transactions, total, nil
}

// Update 更新交易
func (r *transactionRepository) Update(transaction *models.Transaction) error {
	if transaction.ID == 0 {
		return errors.New("transaction id cannot be zero")
	}

	if err := database.DB.Save(transaction).Error; err != nil {
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	return nil
}

// Delete 删除交易（单条）
func (r *transactionRepository) Delete(id int64) error {
	if id == 0 {
		return errors.New("transaction id cannot be zero")
	}

	if err := database.DB.Where("id = ?", id).Delete(&models.Transaction{}).Error; err != nil {
		return fmt.Errorf("failed to delete transaction: %w", err)
	}

	return nil
}

// DeleteBatch 批量删除交易
func (r *transactionRepository) DeleteBatch(ids []int64) error {
	if len(ids) == 0 {
		return errors.New("ids list cannot be empty")
	}

	if err := database.DB.Where("id IN ?", ids).Delete(&models.Transaction{}).Error; err != nil {
		return fmt.Errorf("failed to delete transactions in batch: %w", err)
	}

	return nil
}

// GetTotalBalance 获取交易统计（收入、支出、净额等）
func (r *transactionRepository) GetTotalBalance(userID int64, filters *request.ListTransactionRequest) (*models.Transaction, error) {
	query := database.DB.Where("user_id = ?", userID)

	if filters.Type != "" {
		query = query.Where("type = ?", filters.Type)
	}

	if filters.CategoryID != nil {
		query = query.Where("category_id = ?", *filters.CategoryID)
	}

	if filters.AccountID != nil {
		query = query.Where("account_id = ?", *filters.AccountID)
	}

	if !filters.StartDate.IsZero() {
		query = query.Where("DATE(transaction_date) >= ?", filters.StartDate.Format("2006-01-02"))
	}

	if !filters.EndDate.IsZero() {
		query = query.Where("DATE(transaction_date) <= ?", filters.EndDate.Format("2006-01-02"))
	}

	// 注意：这里返回的不是真实的Transaction对象，而是统计结果
	// 实际项目中应该返回专门的统计响应结构
	var transaction models.Transaction
	err := query.
		Select("COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) as amount").
		First(&transaction).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to get total balance: %w", err)
	}

	return &transaction, nil
}

// GetMonthlyStatistics 获取月度统计
func (r *transactionRepository) GetMonthlyStatistics(userID int64, month time.Time) (*models.Transaction, error) {
	startDate := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
	endDate := startDate.AddDate(0, 1, -1).Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	var transaction models.Transaction
	err := database.DB.
		Where("user_id = ? AND transaction_date >= ? AND transaction_date <= ?", userID, startDate, endDate).
		Select("COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) as amount").
		First(&transaction).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to get monthly statistics: %w", err)
	}

	return &transaction, nil
}

// GetCategoryStatistics 获取分类统计
func (r *transactionRepository) GetCategoryStatistics(userID int64, startDate, endDate time.Time) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := database.DB.
		Where("user_id = ? AND transaction_date >= ? AND transaction_date <= ?", userID, startDate, endDate).
		Group("category_id").
		Select("category_id, SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END) as amount, COUNT(*) as id").
		Find(&transactions).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get category statistics: %w", err)
	}

	return transactions, nil
}

// GetDateRangeStatistics 获取日期范围内的统计
func (r *transactionRepository) GetDateRangeStatistics(userID int64, startDate, endDate time.Time) (*models.Transaction, error) {
	var transaction models.Transaction
	err := database.DB.
		Where("user_id = ? AND transaction_date >= ? AND transaction_date <= ?", userID, startDate, endDate).
		Select(
			"COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) as amount, "+
				"COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0) as id",
		).
		First(&transaction).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to get date range statistics: %w", err)
	}

	return &transaction, nil
}
