package request

// SendVerificationCodeRequest 发送验证码请求
type SendVerificationCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username         string `json:"username" binding:"required,min=3,max=50"`
	Email            string `json:"email" binding:"required,email"`
	Phone            string `json:"phone" binding:"omitempty,min=11,max=20"`
	Password         string `json:"password" binding:"required,min=6,max=32"`
	DisplayName      string `json:"displayName" binding:"omitempty,max=50"`
	VerificationCode string `json:"verificationCode" binding:"required,len=6"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RefreshTokenRequest 刷新Token请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// UpdateUserRequest 更新用户信息请求
type UpdateUserRequest struct {
	DisplayName *string `json:"displayName" binding:"omitempty,max=50"`
	Phone       *string `json:"phone" binding:"omitempty,min=11,max=20"`
	Currency    *string `json:"currency" binding:"omitempty,len=3"`
	Theme       *string `json:"theme" binding:"omitempty,oneof=light dark"`
	Language    *string `json:"language" binding:"omitempty"`
	DarkMode    *bool   `json:"darkMode"`
}

// UpdatePasswordRequest 修改密码请求
type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6,max=32"`
}
