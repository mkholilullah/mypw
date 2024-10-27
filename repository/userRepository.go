package repository

import (
	"github.com/roamercodes/mypw/domain"
	"gorm.io/gorm"
)
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository membuat instance UserRepositoryImpl dengan GORM
func NewUserRepository(DB *gorm.DB) domain.UserRepository {
	return &UserRepository{DB}
}

func (r *UserRepository) GetById(id int) (*domain.User, error) {
    var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) Create(user *domain.User) error {
    return r.DB.Create(user).Error
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByUserName(username string) (*domain.User, error) {
	var user domain.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}