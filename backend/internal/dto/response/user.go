package response

import "time"

// AuthResponse 认证响应
type AuthResponse struct {
	UserID       int64  `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	DisplayName  string `json:"display_name,omitempty"`
	AvatarURL    string `json:"avatar_url,omitempty"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

// UserResponse 用户信息响应
type UserResponse struct {
	ID              int64      `json:"id"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Phone           string     `json:"phone,omitempty"`
	DisplayName     string     `json:"display_name,omitempty"`
	AvatarURL       string     `json:"avatar_url,omitempty"`
	Verified        bool       `json:"verified"`
	Currency        string     `json:"currency"`
	Theme           string     `json:"theme"`
	Language        string     `json:"language"`
	DarkMode        bool       `json:"dark_mode"`
	GestureLock     bool       `json:"gesture_lock"`
	ContinuousDays  int        `json:"continuous_days"`
	TotalRecords    int        `json:"total_records"`
	TotalBadges     int        `json:"total_badges"`
	MembershipLevel string     `json:"membership_level"`
	CreatedAt       time.Time  `json:"created_at"`
	LastLoginAt     *time.Time `json:"last_login_at,omitempty"`
}

// UserStatsResponse 用户统计响应
type UserStatsResponse struct {
	TotalAssets     float64 `json:"total_assets"`
	TotalDebt       float64 `json:"total_debt"`
	NetWorth        float64 `json:"net_worth"`
	TotalRecords    int     `json:"total_records"`
	ContinuousDays  int     `json:"continuous_days"`
	MonthIncome     float64 `json:"month_income"`
	MonthExpense    float64 `json:"month_expense"`
	MonthNet        float64 `json:"month_net"`
	ActiveBudgets   int     `json:"active_budgets"`
	ActiveBills     int     `json:"active_bills"`
	ActiveSavings   int     `json:"active_savings"`
	ActiveWishlists int     `json:"active_wishlists"`
}

// TokenResponse Token响应
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
