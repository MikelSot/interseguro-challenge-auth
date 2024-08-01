package token

import "github.com/MikelSot/interseguro-challenge-auth/model"

type UseCase interface {
	Generate(m model.User) (string, error)
}
