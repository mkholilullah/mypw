package usecase

import (
	"github.com/roamercodes/mypw/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase adalah implementasi dari logika bisnis
// Usecase || Service
type UserUsecase struct {
	UserRepository domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: repo,
	}
}

func (usr *UserUsecase) GetUserById(id int) (*domain.User, error) {
	return usr.UserRepository.GetById(id)
}

// func (usr *UserUsecase) CreateUser(user *domain.User) error {
// 	return usr.UserRepository.Create(user)
// }

func (uc *UserUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	// hash pw befored saved
	hasedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hasedPass)
	err = uc.UserRepository.Create(user)
	return user, err
}