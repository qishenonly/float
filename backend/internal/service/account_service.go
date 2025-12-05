package service

import (
	"errors"

	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/dto/response"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/internal/repository"
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
	accounts, err := s.accountRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var result []*response.AccountResponse
	for _, acc := range accounts {
		result = append(result, s.toAccountResponse(acc))
	}
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
		return nil, errors.New("创建账户失败")
	}

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
	// 注意：余额通常不通过此接口直接修改，而是通过交易记录变动
	// 但如果需要修正初始余额，可以在这里处理，同时需要考虑对当前余额的影响
	// 简单起见，这里暂不处理余额修正逻辑，或者假设Update仅修改属性

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
