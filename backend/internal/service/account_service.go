package service

import (
	"errors"
	"fmt"

	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/dto/response"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/internal/repository"
	"github.com/qiuhaonan/float-backend/pkg/logger"
)

// AccountService 账户服务接口
type AccountService interface {
	GetAccounts(userID int64) ([]*response.AccountResponse, error)
	GetAccountByID(userID int64, accountID int64) (*response.AccountResponse, error)
	CreateAccount(userID int64, req *request.CreateAccountRequest) (*response.AccountResponse, error)
	UpdateAccount(userID int64, accountID int64, req *request.UpdateAccountRequest) error
	DeleteAccount(userID int64, accountID int64) error
	GetAccountBalance(userID int64) (*response.AccountBalanceResponse, error)
}

type accountService struct {
	accountRepo repository.AccountRepository
}

// NewAccountService 创建账户服务实例
func NewAccountService() AccountService {
	return &accountService{
		accountRepo: repository.NewAccountRepository(),
	}
}

// GetAccounts 获取账户列表
func (s *accountService) GetAccounts(userID int64) ([]*response.AccountResponse, error) {
	timer := logger.NewTimer("获取账户列表")
	logger.Info(fmt.Sprintf("[Service][账户] 获取账户列表 | 用户ID: %d", userID))

	accounts, err := s.accountRepo.FindByUserID(userID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][账户] 查询失败 | 用户ID: %d | 错误: %v", userID, err))
		timer.LogError("查询账户列表失败")
		return nil, err
	}

	var result []*response.AccountResponse
	for _, acc := range accounts {
		result = append(result, s.toAccountResponse(acc))
	}

	logger.Info(fmt.Sprintf("[Service][账户] 获取成功 | 用户ID: %d | 账户数: %d", userID, len(result)))
	timer.LogSlowWithThreshold("获取账户列表完成", 300)
	return result, nil
}

// GetAccountByID 获取账户详情
func (s *accountService) GetAccountByID(userID int64, accountID int64) (*response.AccountResponse, error) {
	account, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		return nil, err
	}

	if account.UserID != userID {
		return nil, errors.New("无权访问该账户")
	}

	return s.toAccountResponse(account), nil
}

// CreateAccount 创建账户
func (s *accountService) CreateAccount(userID int64, req *request.CreateAccountRequest) (*response.AccountResponse, error) {
	timer := logger.NewTimer("创建账户")
	logger.Info(fmt.Sprintf("[Service][账户] 创建账户 | 用户ID: %d | 账户名: %s | 初始余额: %.2f", userID, req.AccountName, req.InitialBalance))

	includeInTotal := true
	if req.IncludeInTotal != nil {
		includeInTotal = *req.IncludeInTotal
	}

	account := &models.Account{
		UserID:         userID,
		AccountType:    req.AccountType,
		AccountName:    req.AccountName,
		AccountNumber:  req.AccountNumber,
		Icon:           req.Icon,
		Color:          req.Color,
		Balance:        req.InitialBalance, // 初始余额即为当前余额
		InitialBalance: req.InitialBalance,
		IncludeInTotal: includeInTotal,
		DisplayOrder:   req.DisplayOrder,
		IsActive:       true,
	}

	if err := s.accountRepo.Create(account); err != nil {
		logger.Error(fmt.Sprintf("[Service][账户] 创建失败 | 用户ID: %d | 账户名: %s | 错误: %v", userID, req.AccountName, err))
		timer.LogError("创建账户失败")
		return nil, errors.New("创建账户失败")
	}

	logger.Info(fmt.Sprintf("[Service][账户] 创建成功 | 用户ID: %d | 账户ID: %d", userID, account.ID))
	timer.LogSlowWithThreshold("创建账户完成", 500)
	return s.toAccountResponse(account), nil
}

// UpdateAccount 更新账户
func (s *accountService) UpdateAccount(userID int64, accountID int64, req *request.UpdateAccountRequest) error {
	account, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		return err
	}

	if account.UserID != userID {
		return errors.New("无权修改该账户")
	}

	if req.AccountName != nil {
		account.AccountName = *req.AccountName
	}
	if req.AccountNumber != nil {
		account.AccountNumber = *req.AccountNumber
	}
	if req.Icon != nil {
		account.Icon = *req.Icon
	}
	if req.Color != nil {
		account.Color = *req.Color
	}
	if req.IncludeInTotal != nil {
		account.IncludeInTotal = *req.IncludeInTotal
	}
	if req.DisplayOrder != nil {
		account.DisplayOrder = *req.DisplayOrder
	}
	if req.IsActive != nil {
		account.IsActive = *req.IsActive
	}

	// 如果请求中包含Balance，说明是用户手动修正余额
	if req.Balance != nil {
		account.Balance = *req.Balance
		logger.Info(fmt.Sprintf("[Service][账户] 手动修正余额 | 用户ID: %d | 账户ID: %d | 新余额: %.2f", userID, accountID, *req.Balance))
	}

	return s.accountRepo.Update(account)
}

// DeleteAccount 删除账户
func (s *accountService) DeleteAccount(userID int64, accountID int64) error {
	account, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		return err
	}

	if account.UserID != userID {
		return errors.New("无权删除该账户")
	}

	return s.accountRepo.Delete(accountID)
}

// GetAccountBalance 获取账户余额汇总
func (s *accountService) GetAccountBalance(userID int64) (*response.AccountBalanceResponse, error) {
	totalBalance, err := s.accountRepo.GetTotalBalance(userID)
	if err != nil {
		return nil, err
	}

	// 目前只实现了资产账户，负债账户（信用卡等）后续实现
	// 暂时将总资产视为净资产
	return &response.AccountBalanceResponse{
		TotalBalance: totalBalance,
		AssetBalance: totalBalance,
		DebtBalance:  0,
	}, nil
}

// toAccountResponse 转换为响应对象
func (s *accountService) toAccountResponse(account *models.Account) *response.AccountResponse {
	return &response.AccountResponse{
		ID:             account.ID,
		AccountType:    account.AccountType,
		AccountName:    account.AccountName,
		AccountNumber:  account.AccountNumber,
		Icon:           account.Icon,
		Color:          account.Color,
		Balance:        account.Balance,
		InitialBalance: account.InitialBalance,
		IncludeInTotal: account.IncludeInTotal,
		DisplayOrder:   account.DisplayOrder,
		IsActive:       account.IsActive,
		CreatedAt:      account.CreatedAt,
		UpdatedAt:      account.UpdatedAt,
	}
}
