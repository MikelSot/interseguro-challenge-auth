package user

import (
	"github.com/AJRDRGZ/db-query-builder/models"

	"github.com/MikelSot/interseguro-challenge-auth/model"
)

type UseCase interface {
	Create(m model.User) (model.User, error)

	GetByEmail(email string) (model.User, error)
}

type Storage interface {
	Create(m model.User) error

	GetWhere(specification models.FieldsSpecification) (model.User, error)
}
