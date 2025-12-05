package repository

import (
	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/pkg/database"
)

// AccountRepository 账户仓库接口
type AccountRepository interface {
	Create(account *models.Account) error
	FindByID(id int64) (*models.Account, error)
	FindByUserID(userID int64) ([]*models.Account, error)
	Update(account *models.Account) error
	Delete(id int64) error
	GetTotalBalance(userID int64) (float64, error)
}

type accountRepository struct{}

// NewAccountRepository 创建账户仓库实例
func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}

// Create 创建账户
func (r *accountRepository) Create(account *models.Account) error {
	return database.DB.Create(account).Error
}

// FindByID 根据ID查找账户
func (r *accountRepository) FindByID(id int64) (*models.Account, error) {
	var account models.Account
	err := database.DB.Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// FindByUserID 根据用户ID查找账户列表
func (r *accountRepository) FindByUserID(userID int64) ([]*models.Account, error) {
	var accounts []*models.Account
	err := database.DB.Where("user_id = ? AND is_active = ?", userID, true).
		Order("display_order ASC, created_at ASC").
		Find(&accounts).Error
	return accounts, err
}

// Update 更新账户
func (r *accountRepository) Update(account *models.Account) error {
	return database.DB.Save(account).Error
}

// Delete 删除账户（软删除）
func (r *accountRepository) Delete(id int64) error {
	// 实际项目中可能使用软删除字段，这里根据需求使用GORM的Delete
	// 如果模型中有DeletedAt字段，GORM会自动进行软删除
	// 这里我们使用IsActive字段进行逻辑删除
	return database.DB.Model(&models.Account{}).Where("id = ?", id).Update("is_active", false).Error
}

// GetTotalBalance 获取用户总资产（仅统计include_in_total=true的账户）
func (r *accountRepository) GetTotalBalance(userID int64) (float64, error) {
	var total float64
	err := database.DB.Model(&models.Account{}).
		Where("user_id = ? AND is_active = ? AND include_in_total = ?", userID, true, true).
		Select("COALESCE(SUM(balance), 0)").
		Scan(&total).Error
	return total, err
}
