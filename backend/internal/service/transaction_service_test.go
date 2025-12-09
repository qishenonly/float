package service

import (
	"errors"
	"testing"
	"time"

	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTransactionRepository Mock TransactionRepository
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) CreateBatch(transactions []*models.Transaction) error {
	args := m.Called(transactions)
	return args.Error(0)
}

func (m *MockTransactionRepository) FindByID(id int64) (*models.Transaction, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) FindByUserID(userID int64, filters *request.ListTransactionRequest) ([]*models.Transaction, int64, error) {
	args := m.Called(userID, filters)
	return args.Get(0).([]*models.Transaction), args.Get(1).(int64), args.Error(2)
}

func (m *MockTransactionRepository) Update(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTransactionRepository) DeleteBatch(ids []int64) error {
	args := m.Called(ids)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetTotalBalance(userID int64, filters *request.ListTransactionRequest) (*models.Transaction, error) {
	args := m.Called(userID, filters)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetMonthlyStatistics(userID int64, month time.Time) (*models.Transaction, error) {
	args := m.Called(userID, month)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetCategoryStatistics(userID int64, startDate, endDate time.Time) ([]*models.Transaction, error) {
	args := m.Called(userID, startDate, endDate)
	return args.Get(0).([]*models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetDateRangeStatistics(userID int64, startDate, endDate time.Time) (*models.Transaction, error) {
	args := m.Called(userID, startDate, endDate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Transaction), args.Error(1)
}

// MockAccountRepository Mock AccountRepository
type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) Create(account *models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountRepository) FindByID(id int64) (*models.Account, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Account), args.Error(1)
}

func (m *MockAccountRepository) FindByUserID(userID int64) ([]*models.Account, error) {
	args := m.Called(userID)
	return args.Get(0).([]*models.Account), args.Error(1)
}

func (m *MockAccountRepository) Update(account *models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAccountRepository) GetTotalBalance(userID int64) (float64, error) {
	args := m.Called(userID)
	return args.Get(0).(float64), args.Error(1)
}

// MockCategoryRepository Mock CategoryRepository
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(category *models.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) FindByID(id int64) (*models.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindByUserID(userID int64, categoryType string) ([]*models.Category, error) {
	args := m.Called(userID, categoryType)
	return args.Get(0).([]*models.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindAll(userID int64) ([]*models.Category, error) {
	args := m.Called(userID)
	return args.Get(0).([]*models.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(category *models.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetSystemCategories(categoryType string) ([]*models.Category, error) {
	args := m.Called(categoryType)
	return args.Get(0).([]*models.Category), args.Error(1)
}

// TestCreateTransaction 测试创建交易
func TestCreateTransaction(t *testing.T) {
	mockTransRepo := new(MockTransactionRepository)
	mockAccountRepo := new(MockAccountRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	service := NewTransactionService(mockTransRepo, mockAccountRepo, mockCategoryRepo)

	userID := int64(1)
	accountID := int64(100)
	categoryID := int64(10)
	transactionDate := time.Now()

	req := &request.CreateTransactionRequest{
		Type:            "expense",
		CategoryID:      &categoryID,
		AccountID:       &accountID,
		Amount:          100.00,
		Currency:        "CNY",
		Title:           "测试交易",
		TransactionDate: transactionDate,
	}

	// Mock账户查询
	account := &models.Account{
		ID:     accountID,
		UserID: userID,
	}
	mockAccountRepo.On("FindByID", accountID).Return(account, nil)

	// Mock分类查询
	category := &models.Category{
		ID:     categoryID,
		UserID: userID,
		Type:   "expense",
	}
	mockCategoryRepo.On("FindByID", categoryID).Return(category, nil)

	// Mock交易创建
	mockTransRepo.On("Create", mock.MatchedBy(func(t *models.Transaction) bool {
		return t.UserID == userID && t.Amount == 100.00
	})).Run(func(args mock.Arguments) {
		trans := args.Get(0).(*models.Transaction)
		trans.ID = 1
	}).Return(nil)

	// Mock账户更新
	mockAccountRepo.On("Update", mock.MatchedBy(func(a *models.Account) bool {
		return a.ID == accountID
	})).Return(nil)

	// Mock交易查询
	mockTransRepo.On("FindByID", int64(1)).Return(&models.Transaction{
		ID:       1,
		UserID:   userID,
		Type:     "expense",
		Amount:   100.00,
		Category: category,
		Account:  account,
	}, nil)

	// 执行
	resp, err := service.CreateTransaction(userID, req)

	// 断言
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, userID, resp.UserID)
	assert.Equal(t, 100.00, resp.Amount)
	assert.Equal(t, "expense", resp.Type)
}

// TestUpdateTransaction 测试更新交易
func TestUpdateTransaction(t *testing.T) {
	mockTransRepo := new(MockTransactionRepository)
	mockAccountRepo := new(MockAccountRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	service := NewTransactionService(mockTransRepo, mockAccountRepo, mockCategoryRepo)

	userID := int64(1)
	transactionID := int64(1)

	oldTransaction := &models.Transaction{
		ID:     transactionID,
		UserID: userID,
		Amount: 100.00,
		Type:   "expense",
	}

	req := &request.UpdateTransactionRequest{
		Amount: 150.00,
	}

	mockTransRepo.On("FindByID", transactionID).Return(oldTransaction, nil)
	mockTransRepo.On("Update", mock.MatchedBy(func(t *models.Transaction) bool {
		return t.ID == transactionID && t.Amount == 150.00
	})).Return(nil)
	mockTransRepo.On("FindByID", transactionID).Return(&models.Transaction{
		ID:     transactionID,
		UserID: userID,
		Amount: 150.00,
		Type:   "expense",
	}, nil)

	resp, err := service.UpdateTransaction(userID, transactionID, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 150.00, resp.Amount)
}

// TestDeleteTransaction 测试删除交易
func TestDeleteTransaction(t *testing.T) {
	mockTransRepo := new(MockTransactionRepository)
	mockAccountRepo := new(MockAccountRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	service := NewTransactionService(mockTransRepo, mockAccountRepo, mockCategoryRepo)

	userID := int64(1)
	transactionID := int64(1)
	accountID := int64(100)

	transaction := &models.Transaction{
		ID:        transactionID,
		UserID:    userID,
		Type:      "expense",
		Amount:    100.00,
		AccountID: &accountID,
	}

	account := &models.Account{
		ID:      accountID,
		UserID:  userID,
		Balance: 500.00,
	}

	mockTransRepo.On("FindByID", transactionID).Return(transaction, nil)
	mockAccountRepo.On("FindByID", accountID).Return(account, nil)
	mockAccountRepo.On("Update", mock.MatchedBy(func(a *models.Account) bool {
		return a.Balance == 600.00 // 恢复金额
	})).Return(nil)
	mockTransRepo.On("Delete", transactionID).Return(nil)

	err := service.DeleteTransaction(userID, transactionID)

	assert.NoError(t, err)
	mockTransRepo.AssertCalled(t, "Delete", transactionID)
}

// TestListTransactions 测试查询交易列表
func TestListTransactions(t *testing.T) {
	mockTransRepo := new(MockTransactionRepository)
	mockAccountRepo := new(MockAccountRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	service := NewTransactionService(mockTransRepo, mockAccountRepo, mockCategoryRepo)

	userID := int64(1)
	filters := &request.ListTransactionRequest{
		Page:     1,
		PageSize: 20,
	}

	transactions := []*models.Transaction{
		{
			ID:     1,
			UserID: userID,
			Amount: 100.00,
			Type:   "expense",
		},
		{
			ID:     2,
			UserID: userID,
			Amount: 50.00,
			Type:   "income",
		},
	}

	mockTransRepo.On("FindByUserID", userID, filters).Return(transactions, int64(2), nil)

	resp, err := service.ListTransactions(userID, filters)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(2), resp.Total)
	assert.Equal(t, 2, len(resp.Items))
}

// TestGetTransactionStatistics 测试获取交易统计
func TestGetTransactionStatistics(t *testing.T) {
	mockTransRepo := new(MockTransactionRepository)
	mockAccountRepo := new(MockAccountRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	service := NewTransactionService(mockTransRepo, mockAccountRepo, mockCategoryRepo)

	userID := int64(1)
	startDate := time.Now().AddDate(0, -1, 0)
	endDate := time.Now()

	transactions := []*models.Transaction{
		{
			ID:     1,
			UserID: userID,
			Amount: 100.00,
			Type:   "income",
		},
		{
			ID:     2,
			UserID: userID,
			Amount: 50.00,
			Type:   "expense",
		},
	}

	mockTransRepo.On("FindByUserID", userID, mock.MatchedBy(func(f *request.ListTransactionRequest) bool {
		return f.StartDate.Equal(startDate) && f.EndDate.Equal(endDate)
	})).Return(transactions, int64(2), nil)

	resp, err := service.GetTransactionStatistics(userID, startDate, endDate)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 100.00, resp.TotalIncome)
	assert.Equal(t, 50.00, resp.TotalExpense)
	assert.Equal(t, 50.00, resp.NetAmount)
	assert.Equal(t, int64(2), resp.TransactionCnt)
}

// TestCreateBatchTransactions 测试批量创建交易
func TestCreateBatchTransactions(t *testing.T) {
	mockTransRepo := new(MockTransactionRepository)
	mockAccountRepo := new(MockAccountRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	service := NewTransactionService(mockTransRepo, mockAccountRepo, mockCategoryRepo)

	userID := int64(1)
	accountID := int64(100)

	req := &request.BulkCreateTransactionRequest{
		Transactions: []request.CreateTransactionRequest{
			{
				Type:            "expense",
				AccountID:       &accountID,
				Amount:          100.00,
				TransactionDate: time.Now(),
			},
		},
	}

	account := &models.Account{
		ID:     accountID,
		UserID: userID,
	}

	mockAccountRepo.On("FindByID", accountID).Return(account, nil)
	mockTransRepo.On("Create", mock.Anything).Run(func(args mock.Arguments) {
		trans := args.Get(0).(*models.Transaction)
		trans.ID = 1
	}).Return(nil)
	mockAccountRepo.On("Update", mock.Anything).Return(nil)
	mockTransRepo.On("FindByID", int64(1)).Return(&models.Transaction{
		ID:      1,
		UserID:  userID,
		Amount:  100.00,
		Type:    "expense",
		Account: account,
	}, nil)

	resp, err := service.CreateBatchTransactions(userID, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(1), resp.SuccessCount)
}

// TestInvalidCategoryID 测试无效的分类ID
func TestInvalidCategoryID(t *testing.T) {
	mockTransRepo := new(MockTransactionRepository)
	mockAccountRepo := new(MockAccountRepository)
	mockCategoryRepo := new(MockCategoryRepository)

	service := NewTransactionService(mockTransRepo, mockAccountRepo, mockCategoryRepo)

	userID := int64(1)
	accountID := int64(100)
	invalidCategoryID := int64(999)

	req := &request.CreateTransactionRequest{
		Type:            "expense",
		CategoryID:      &invalidCategoryID,
		AccountID:       &accountID,
		Amount:          100.00,
		TransactionDate: time.Now(),
	}

	account := &models.Account{
		ID:     accountID,
		UserID: userID,
	}

	mockAccountRepo.On("FindByID", accountID).Return(account, nil)
	mockCategoryRepo.On("FindByID", invalidCategoryID).Return(nil, errors.New("category not found"))

	resp, err := service.CreateTransaction(userID, req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "invalid category", err.Error())
}
