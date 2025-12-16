package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/qiuhaonan/float-backend/internal/dto/request"
	"github.com/qiuhaonan/float-backend/internal/dto/response"
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/internal/repository"
	"github.com/qiuhaonan/float-backend/internal/utils"
	"github.com/qiuhaonan/float-backend/pkg/email"
	"github.com/qiuhaonan/float-backend/pkg/logger"
	"github.com/spf13/viper"
)

// UserService 用户服务接口
type UserService interface {
	SendVerificationCode(req *request.SendVerificationCodeRequest) error
	Register(req *request.RegisterRequest) (*response.AuthResponse, error)
	Login(req *request.LoginRequest) (*response.AuthResponse, error)
	RefreshToken(refreshToken string) (*response.TokenResponse, error)
	GetUserByID(userID int64) (*response.UserResponse, error)
	UpdateUser(userID int64, req *request.UpdateUserRequest) error
	UpdatePassword(userID int64, req *request.UpdatePasswordRequest) error
	GetUserStats(userID int64) (*response.UserStatsResponse, error)
	GetSystemOverview() (map[string]interface{}, error)
	ListUsers(page, pageSize int) ([]*response.UserResponse, int64, error)
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

// SendVerificationCode 发送邮箱验证码
func (s *userService) SendVerificationCode(req *request.SendVerificationCodeRequest) error {
	timer := logger.NewTimer("发送验证码")
	sanitizedEmail := logger.SanitizeEmail(req.Email)
	logger.Info(fmt.Sprintf("[Service][验证码] 发送请求 | 邮箱: %s", sanitizedEmail))

	// 注意：这里不检查邮箱是否已注册，仅生成和发送验证码
	// 注册时才进行邮箱重复检查
	// 这样可以快速响应用户请求，提供更好的用户体验

	// 生成验证码
	code, err := email.GenerateVerificationCode()
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][验证码] 生成失败 | 邮箱: %s | 错误: %v", sanitizedEmail, err))
		timer.LogError("生成验证码失败")
		return errors.New("生成验证码失败")
	}

	// 存储验证码到 Redis
	if err := email.StoreVerificationCode(req.Email, code); err != nil {
		logger.Error(fmt.Sprintf("[Service][验证码] 存储失败 | 邮箱: %s | 错误: %v", sanitizedEmail, err))
		timer.LogError("存储验证码失败")
		return errors.New("存储验证码失败")
	}

	// 发送邮件
	emailService := email.GetService()
	if err := emailService.SendVerificationCode(req.Email, code); err != nil {
		logger.Error(fmt.Sprintf("[Service][验证码] 发送失败 | 邮箱: %s | 错误: %v", sanitizedEmail, err))
		timer.LogError("发送验证码邮件失败")
		return errors.New("发送验证码邮件失败")
	}

	logger.Info(fmt.Sprintf("[Service][验证码] 发送成功 | 邮箱: %s", sanitizedEmail))
	timer.LogSlowWithThreshold("发送验证码成功", 500) // 500ms 阈值
	return nil
}

// Register 用户注册
func (s *userService) Register(req *request.RegisterRequest) (*response.AuthResponse, error) {
	timer := logger.NewTimer("用户注册")
	sanitizedEmail := logger.SanitizeEmail(req.Email)
	logger.Info(fmt.Sprintf("[Service][注册] 开始注册 | 邮箱: %s | 用户名: %s", sanitizedEmail, req.Username))

	// 验证验证码是否提供
	if req.VerificationCode == "" {
		logger.Warn(fmt.Sprintf("[Service][注册] 缺少验证码 | 邮箱: %s", sanitizedEmail))
		timer.LogError("缺少验证码")
		return nil, errors.New("验证码不能为空")
	}

	// 验证邮箱验证码
	verified, err := email.VerifyCode(req.Email, req.VerificationCode)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][注册] 验证码验证失败 | 邮箱: %s | 错误: %v", sanitizedEmail, err))
		timer.LogError("验证码验证失败")
		return nil, errors.New("验证码验证失败，请重试")
	}
	if !verified {
		logger.Warn(fmt.Sprintf("[Service][注册] 验证码错误或已过期 | 邮箱: %s", sanitizedEmail))
		timer.LogError("验证码错误或已过期")
		return nil, errors.New("验证码错误或已过期")
	}

	// 检查用户名是否已存在
	existingUser, _ := s.userRepo.FindByUsername(req.Username)
	if existingUser != nil {
		logger.Warn(fmt.Sprintf("[Service][注册] 用户名已存在 | 用户名: %s", req.Username))
		timer.LogError("用户名已存在")
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	existingEmail, _ := s.userRepo.FindByEmail(req.Email)
	if existingEmail != nil {
		logger.Warn(fmt.Sprintf("[Service][注册] 邮箱已被注册 | 邮箱: %s", sanitizedEmail))
		timer.LogError("邮箱已被注册")
		return nil, errors.New("邮箱已被注册")
	}

	// 哈希密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][注册] 密码加密失败 | 邮箱: %s | 错误: %v", sanitizedEmail, err))
		timer.LogError("密码加密失败")
		return nil, errors.New("密码加密失败")
	}

	// 创建用户
	var phone *string
	if req.Phone != "" {
		phone = &req.Phone
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		Phone:        phone,
		PasswordHash: hashedPassword,
		DisplayName:  req.DisplayName,
		Currency:     "CNY",
		Theme:        "light",
		Language:     "zh-CN",
		DarkMode:     false,
		GestureLock:  true,
	}

	if err := s.userRepo.Create(user); err != nil {
		logger.Error(fmt.Sprintf("[Service][注册] 创建用户失败 | 邮箱: %s | 错误: %v", sanitizedEmail, err))
		timer.LogError("创建用户失败")
		return nil, errors.New("创建用户失败")
	}
	logger.Info(fmt.Sprintf("[Service][注册] 用户创建成功 | 用户ID: %d", user.ID))

	// 生成 Token
	accessToken, err := utils.GenerateToken(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][注册] 生成访问令牌失败 | 用户ID: %d | 错误: %v", user.ID, err))
		timer.LogError("生成访问令牌失败")
		return nil, errors.New("生成访问令牌失败")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username, user.Email)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][注册] 生成刷新令牌失败 | 用户ID: %d | 错误: %v", user.ID, err))
		timer.LogError("生成刷新令牌失败")
		return nil, errors.New("生成刷新令牌失败")
	}

	// 发送欢迎邮件
	emailService := email.GetService()
	displayName := user.DisplayName
	if displayName == "" {
		displayName = user.Username
	}
	if err := emailService.SendWelcome(user.Email, displayName); err != nil {
		logger.Error(fmt.Sprintf("[Service][注册] 发送欢迎邮件失败 | 用户ID: %d | 错误: %v", user.ID, err))
		// 邮件发送失败不影响注册结果
	}

	logger.Info(fmt.Sprintf("[Service][注册] 注册成功 | 用户ID: %d | 邮箱: %s | 用户名: %s", user.ID, sanitizedEmail, user.Username))
	timer.LogSlowWithThreshold("注册流程完成", 1000) // 1秒阈值
	return &response.AuthResponse{
		UserID:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		DisplayName:  user.DisplayName,
		AvatarURL:    user.AvatarURL,
		Role:         user.Role,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    viper.GetInt("jwt.access_token_expire"),
	}, nil
}

// Login 用户登录
func (s *userService) Login(req *request.LoginRequest) (*response.AuthResponse, error) {
	timer := logger.NewTimer("用户登录")
	sanitizedEmail := logger.SanitizeEmail(req.Email)
	logger.Info(fmt.Sprintf("[Service][登录] 开始登录 | 邮箱: %s", sanitizedEmail))

	// 查找用户
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		logger.Warn(fmt.Sprintf("[Service][登录] 用户不存在 | 邮箱: %s", sanitizedEmail))
		timer.LogError("用户不存在")
		return nil, errors.New("邮箱或密码错误")
	}

	// 验证密码
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		logger.Warn(fmt.Sprintf("[Service][登录] 密码错误 | 用户ID: %d | 邮箱: %s", user.ID, sanitizedEmail))
		timer.LogError("密码验证失败")
		return nil, errors.New("邮箱或密码错误")
	}

	// 更新最后登录时间
	if err := s.userRepo.UpdateLastLogin(user.ID); err != nil {
		logger.Error(fmt.Sprintf("[Service][登录] 更新登录时间失败 | 用户ID: %d | 错误: %v", user.ID, err))
	}

	// 生成 Token
	accessToken, err := utils.GenerateToken(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][登录] 生成访问令牌失败 | 用户ID: %d | 错误: %v", user.ID, err))
		timer.LogError("生成访问令牌失败")
		return nil, errors.New("生成访问令牌失败")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username, user.Email)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][登录] 生成刷新令牌失败 | 用户ID: %d | 错误: %v", user.ID, err))
		timer.LogError("生成刷新令牌失败")
		return nil, errors.New("生成刷新令牌失败")
	}

	logger.Info(fmt.Sprintf("[Service][登录] 登录成功 | 用户ID: %d | 用户名: %s", user.ID, user.Username))
	timer.LogSlowWithThreshold("登录流程完成", 500) // 500ms 阈值
	return &response.AuthResponse{
		UserID:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		DisplayName:  user.DisplayName,
		AvatarURL:    user.AvatarURL,
		Role:         user.Role,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    viper.GetInt("jwt.access_token_expire"),
	}, nil
}

// RefreshToken 刷新访问令牌
func (s *userService) RefreshToken(refreshToken string) (*response.TokenResponse, error) {
	timer := logger.NewTimer("刷新令牌")
	logger.Info("[Service][刷新令牌] 开始刷新")

	// 解析 Refresh Token
	claims, err := utils.ParseToken(refreshToken)
	if err != nil {
		logger.Warn(fmt.Sprintf("[Service][刷新令牌] 令牌解析失败 | 错误: %v", err))
		timer.LogError("令牌解析失败")
		return nil, errors.New("刷新令牌无效或已过期")
	}

	// 获取最新用户信息以确保 Role 是最新的
	user, err := s.userRepo.FindByID(claims.UserID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][刷新令牌] 查询用户失败 | 用户ID: %d | 错误: %v", claims.UserID, err))
		timer.LogError("查询用户失败")
		return nil, errors.New("刷新令牌失败: 用户不存在")
	}

	// 生成新的 Access Token
	accessToken, err := utils.GenerateToken(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][刷新令牌] 生成访问令牌失败 | 用户ID: %d | 错误: %v", claims.UserID, err))
		timer.LogError("生成访问令牌失败")
		return nil, errors.New("生成访问令牌失败")
	}

	logger.Info(fmt.Sprintf("[Service][刷新令牌] 刷新成功 | 用户ID: %d", claims.UserID))
	timer.LogSlowWithThreshold("令牌刷新完成", 300) // 300ms 阈值
	return &response.TokenResponse{
		AccessToken: accessToken,
		ExpiresIn:   viper.GetInt("jwt.access_token_expire"),
	}, nil
}

// GetUserByID 获取用户信息
func (s *userService) GetUserByID(userID int64) (*response.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		logger.Error(fmt.Sprintf("[用户服务][获取信息] 查询用户失败 | 用户ID: %d | 错误: %v", userID, err))
		return nil, err
	}

	var phone string
	if user.Phone != nil {
		phone = *user.Phone
	}

	logger.Info(fmt.Sprintf("[用户服务][获取信息] 成功获取用户信息 | 用户ID: %d | 用户名: %s", user.ID, user.Username))
	return &response.UserResponse{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		Phone:           phone,
		DisplayName:     user.DisplayName,
		AvatarURL:       user.AvatarURL,
		Verified:        user.Verified,
		Role:            user.Role,
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
		logger.Error(fmt.Sprintf("[用户服务][更新信息] 查询用户失败 | 用户ID: %d | 错误: %v", userID, err))
		return err
	}

	// 更新字段
	if req.DisplayName != nil {
		user.DisplayName = *req.DisplayName
	}
	if req.Phone != nil {
		user.Phone = req.Phone
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

	if err := s.userRepo.Update(user); err != nil {
		logger.Error(fmt.Sprintf("[用户服务][更新信息] 更新用户失败 | 用户ID: %d | 错误: %v", userID, err))
		return err
	}

	logger.Info(fmt.Sprintf("[用户服务][更新信息] 用户信息更新成功 | 用户ID: %d", userID))
	return nil
}

// UpdatePassword 修改密码
func (s *userService) UpdatePassword(userID int64, req *request.UpdatePasswordRequest) error {
	timer := logger.NewTimer("修改密码")
	logger.Info(fmt.Sprintf("[Service][修改密码] 开始修改 | 用户ID: %d", userID))

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][修改密码] 查询用户失败 | 用户ID: %d | 错误: %v", userID, err))
		timer.LogError("查询用户失败")
		return err
	}

	// 验证旧密码
	if !utils.CheckPassword(req.OldPassword, user.PasswordHash) {
		logger.Warn(fmt.Sprintf("[Service][修改密码] 原密码错误 | 用户ID: %d", userID))
		timer.LogError("原密码验证失败")
		return errors.New("原密码错误")
	}

	// 哈希新密码
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service][修改密码] 密码加密失败 | 用户ID: %d | 错误: %v", userID, err))
		timer.LogError("密码加密失败")
		return errors.New("密码加密失败")
	}

	user.PasswordHash = hashedPassword
	if err := s.userRepo.Update(user); err != nil {
		logger.Error(fmt.Sprintf("[Service][修改密码] 更新密码失败 | 用户ID: %d | 错误: %v", userID, err))
		timer.LogError("更新密码失败")
		return err
	}

	logger.Info(fmt.Sprintf("[Service][修改密码] 修改成功 | 用户ID: %d", userID))
	timer.LogSlowWithThreshold("密码修改完成", 800) // 800ms 阈值
	return nil
}

// GetUserStats 获取用户统计信息
func (s *userService) GetUserStats(userID int64) (*response.UserStatsResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		logger.Error(fmt.Sprintf("[用户服务][获取统计] 查询用户失败 | 用户ID: %d | 错误: %v", userID, err))
		return nil, err
	}

	// TODO: 从其他表计算统计数据
	// 这里先返回模拟数据，后续实现其他模块后再补充
	now := time.Now()
	_ = now

	logger.Info(fmt.Sprintf("[用户服务][获取统计] 成功获取用户统计信息 | 用户ID: %d", userID))
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

// GetSystemOverview 获取系统概览（管理员）
func (s *userService) GetSystemOverview() (map[string]interface{}, error) {
	totalUsers, err := s.userRepo.Count()
	if err != nil {
		return nil, err
	}

	// TODO: Get total transactions count from transaction repo if available

	return map[string]interface{}{
		"total_users":   totalUsers,
		"system_status": "normal",
	}, nil
}

// ListUsers 获取用户列表（管理员）
func (s *userService) ListUsers(page, pageSize int) ([]*response.UserResponse, int64, error) {
	users, err := s.userRepo.FindAll(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.userRepo.Count()
	if err != nil {
		return nil, 0, err
	}

	var userResps []*response.UserResponse
	for _, u := range users {
		var phone string
		if u.Phone != nil {
			phone = *u.Phone
		}
		userResps = append(userResps, &response.UserResponse{
			ID:              u.ID,
			Username:        u.Username,
			Email:           u.Email,
			Phone:           phone,
			DisplayName:     u.DisplayName,
			AvatarURL:       u.AvatarURL,
			Verified:        u.Verified,
			Role:            u.Role,
			CreatedAt:       u.CreatedAt,
			LastLoginAt:     u.LastLoginAt,
			MembershipLevel: u.MembershipLevel,
		})
	}

	return userResps, total, nil
}
