package usecase

import (
	"errors"

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

func (uc *UserUsecase) GetUserById(id int) (*domain.User, error) {
	return uc.UserRepository.GetById(id)
}

// func (usr *UserUsecase) CreateUser(user *domain.User) error {
// 	return usr.UserRepository.Create(user)
// }

func (uc *UserUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	// hash pw befored saved
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	err = uc.UserRepository.Create(user)
	return user, err
}

func (uc *UserUsecase) LoginUser(username, password string) (*domain.User, error) {
	// search user by username
	user, err := uc.UserRepository.GetByUserName(username)
	// user, err := uc.UserRepository.G
	if err != nil {
		return nil, errors.New("username atau password salah!")
	}

	return user, nil
}