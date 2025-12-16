package repository

import (
	"errors"

	"github.com/qiuhaonan/float-backend/internal/models"
	"github.com/qiuhaonan/float-backend/pkg/database"
	"gorm.io/gorm"
)

// UserRepository 用户仓库接口
type UserRepository interface {
	Create(user *models.User) error
	FindByID(id int64) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	UpdateLastLogin(userID int64) error
	Count() (int64, error)
	FindAll(page, pageSize int) ([]*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓库实例
func NewUserRepository() UserRepository {
	return &userRepository{
		db: database.GetDB(),
	}
}

// Create 创建用户
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByID 根据ID查找用户
func (r *userRepository) FindByID(id int64) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// FindByUsername 根据用户名查找用户
func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail 根据邮箱查找用户
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// UpdateLastLogin 更新最后登录时间
func (r *userRepository) UpdateLastLogin(userID int64) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).
		Update("last_login_at", gorm.Expr("NOW()")).Error
}

// Count 统计用户总数
func (r *userRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Count(&count).Error
	return count, err
}

// FindAll 获取用户列表（分页）
func (r *userRepository) FindAll(page, pageSize int) ([]*models.User, error) {
	var users []*models.User
	offset := (page - 1) * pageSize
	err := r.db.Offset(offset).Limit(pageSize).Order("id desc").Find(&users).Error
	return users, err
}
