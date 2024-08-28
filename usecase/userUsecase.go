package usecase

import "github.com/roamercodes/mypw/domain"

// UserUseCase adalah implementasi dari logika bisnis
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

func (usr *UserUsecase) CreateUser(user *domain.User) error {
	return usr.UserRepository.Create(user)
}