package request

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username    string `json:"username" binding:"required,min=3,max=50"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone" binding:"omitempty,min=11,max=20"`
	Password    string `json:"password" binding:"required,min=6,max=32"`
	DisplayName string `json:"display_name" binding:"omitempty,max=50"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RefreshTokenRequest 刷新Token请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UpdateUserRequest 更新用户信息请求
type UpdateUserRequest struct {
	DisplayName *string `json:"display_name" binding:"omitempty,max=50"`
	Phone       *string `json:"phone" binding:"omitempty,min=11,max=20"`
	Currency    *string `json:"currency" binding:"omitempty,len=3"`
	Theme       *string `json:"theme" binding:"omitempty,oneof=light dark"`
	Language    *string `json:"language" binding:"omitempty"`
	DarkMode    *bool   `json:"dark_mode"`
}

// UpdatePasswordRequest 修改密码请求
type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=32"`
}
