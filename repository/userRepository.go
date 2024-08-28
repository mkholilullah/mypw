package repository

import (
	"github.com/roamercodes/mypw/domain"
	"gorm.io/gorm"
)
type UserRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepository membuat instance UserRepositoryImpl dengan GORM
func NewUserRepository(DB *gorm.DB) domain.UserRepository {
	return &UserRepositoryImpl{DB}
}

// Implementasi GetByID
func (r *UserRepositoryImpl) GetById(id int) (*domain.User, error) {
    var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// Implementasi Create menggunakan GORM
func (r *UserRepositoryImpl) Create(user *domain.User) error {
    return r.DB.Create(user).Error
}

// func NewUserRepository(db *sql.DB) domain.UserRepository {
//     return &UserRepositoryI{DB: db}
// }