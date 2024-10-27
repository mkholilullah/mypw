package usecase

import "github.com/roamercodes/mypw/domain"

type authenticationUsecase struct {
	UserRepository domain.UserRepository
}

