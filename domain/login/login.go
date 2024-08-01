package login

import "github.com/MikelSot/interseguro-challenge-auth/model"

type UseCase interface {
	Login(m model.User) (interface{}, error)
}

type UserUseCase interface {
	GetByEmail(email string) (model.User, error)
}

type TokenUseCase interface {
	Generate(m model.User) (string, error)
}
