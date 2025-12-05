package service

import (
	"errors"
	"time"

	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/dto/response"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/internal/repository"
	"github.com/qiuhaonan/float-backend/internal/utils"
	"github.com/spf13/viper"
)

// UserService 用户服务接口
type UserService interface {
	Register(req *request.RegisterRequest) (*response.AuthResponse, error)
	Login(req *request.LoginRequest) (*response.AuthResponse, error)
	RefreshToken(refreshToken string) (*response.TokenResponse, error)
	GetUserByID(userID int64) (*response.UserResponse, error)
	UpdateUser(userID int64, req *request.UpdateUserRequest) error
	UpdatePassword(userID int64, req *request.UpdatePasswordRequest) error
	GetUserStats(userID int64) (*response.UserStatsResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService() UserService {
	return &userService{
		userRepo: repository.NewUserRepository(),
	}
}

// Register 用户注册
func (s *userService) Register(req *request.RegisterRequest) (*response.AuthResponse, error) {
	// 检查用户名是否已存在
	existingUser, _ := s.userRepo.FindByUsername(req.Username)
	if existingUser != nil {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	existingEmail, _ := s.userRepo.FindByEmail(req.Email)
	if existingEmail != nil {
		return nil, errors.New("邮箱已被注册")
	}

	// 哈希密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	// 创建用户
	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: hashedPassword,
		DisplayName:  req.DisplayName,
		Currency:     "CNY",
		Theme:        "light",
		Language:     "zh-CN",
		DarkMode:     false,
		GestureLock:  true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("创建用户失败")
	}

	// 生成 Token
	accessToken, err := utils.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		return nil, errors.New("生成访问令牌失败")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username, user.Email)
	if err != nil {
		return nil, errors.New("生成刷新令牌失败")
	}

	return &response.AuthResponse{
		UserID:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		DisplayName:  user.DisplayName,
		AvatarURL:    user.AvatarURL,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    viper.GetInt("jwt.access_token_expire"),
	}, nil
}

// Login 用户登录
func (s *userService) Login(req *request.LoginRequest) (*response.AuthResponse, error) {
	// 查找用户
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("用户名或密码错误")
	}

	// 更新最后登录时间
	_ = s.userRepo.UpdateLastLogin(user.ID)

	// 生成 Token
	accessToken, err := utils.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		return nil, errors.New("生成访问令牌失败")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username, user.Email)
	if err != nil {
		return nil, errors.New("生成刷新令牌失败")
	}

	return &response.AuthResponse{
		UserID:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		DisplayName:  user.DisplayName,
		AvatarURL:    user.AvatarURL,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    viper.GetInt("jwt.access_token_expire"),
	}, nil
}

// RefreshToken 刷新访问令牌
func (s *userService) RefreshToken(refreshToken string) (*response.TokenResponse, error) {
	// 解析 Refresh Token
	claims, err := utils.ParseToken(refreshToken)
	if err != nil {
		return nil, errors.New("刷新令牌无效或已过期")
	}

	// 生成新的 Access Token
	accessToken, err := utils.GenerateToken(claims.UserID, claims.Username, claims.Email)
	if err != nil {
		return nil, errors.New("生成访问令牌失败")
	}

	return &response.TokenResponse{
		AccessToken: accessToken,
		ExpiresIn:   viper.GetInt("jwt.access_token_expire"),
	}, nil
}

// GetUserByID 获取用户信息
func (s *userService) GetUserByID(userID int64) (*response.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		Phone:           user.Phone,
		DisplayName:     user.DisplayName,
		AvatarURL:       user.AvatarURL,
		Verified:        user.Verified,
		Currency:        user.Currency,
		Theme:           user.Theme,
		Language:        user.Language,
		DarkMode:        user.DarkMode,
		GestureLock:     user.GestureLock,
		ContinuousDays:  user.ContinuousDays,
		TotalRecords:    user.TotalRecords,
		TotalBadges:     user.TotalBadges,
		MembershipLevel: user.MembershipLevel,
		CreatedAt:       user.CreatedAt,
		LastLoginAt:     user.LastLoginAt,
	}, nil
}

// UpdateUser 更新用户信息
func (s *userService) UpdateUser(userID int64, req *request.UpdateUserRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	// 更新字段
	if req.DisplayName != nil {
		user.DisplayName = *req.DisplayName
	}
	if req.Phone != nil {
		user.Phone = *req.Phone
	}
	if req.Currency != nil {
		user.Currency = *req.Currency
	}
	if req.Theme != nil {
		user.Theme = *req.Theme
	}
	if req.Language != nil {
		user.Language = *req.Language
	}
	if req.DarkMode != nil {
		user.DarkMode = *req.DarkMode
	}

	return s.userRepo.Update(user)
}

// UpdatePassword 修改密码
func (s *userService) UpdatePassword(userID int64, req *request.UpdatePasswordRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	// 验证旧密码
	if !utils.CheckPassword(req.OldPassword, user.PasswordHash) {
		return errors.New("原密码错误")
	}

	// 哈希新密码
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("密码加密失败")
	}

	user.PasswordHash = hashedPassword
	return s.userRepo.Update(user)
}

// GetUserStats 获取用户统计信息
func (s *userService) GetUserStats(userID int64) (*response.UserStatsResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	// TODO: 从其他表计算统计数据
	// 这里先返回模拟数据，后续实现其他模块后再补充
	now := time.Now()
	_ = now

	return &response.UserStatsResponse{
		TotalAssets:     0,
		TotalDebt:       0,
		NetWorth:        0,
		TotalRecords:    user.TotalRecords,
		ContinuousDays:  user.ContinuousDays,
		MonthIncome:     0,
		MonthExpense:    0,
		MonthNet:        0,
		ActiveBudgets:   0,
		ActiveBills:     0,
		ActiveSavings:   0,
		ActiveWishlists: 0,
	}, nil
}
