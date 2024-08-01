package user

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/AJRDRGZ/db-query-builder/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/MikelSot/interseguro-challenge-auth/model"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const _defaultMinLenPassword = 5

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{s}
}

func (u User) Create(m model.User) (model.User, error) {
	if err := model.ValidateStructNil(m); err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	m.Email = strings.ToLower(m.Email)

	customError := model.NewError()
	if !emailRegex.MatchString(m.Email) {
		customError.SetCode(model.InvalidEmail)
		customError.SetAPIMessage("¡Upps! el email no es válido")
		customError.SetStatusHTTP(http.StatusBadRequest)

		return model.User{}, customError
	}

	if len(m.Password) < _defaultMinLenPassword {
		customError.SetCode(model.InvalidPassword)
		customError.SetAPIMessage("¡Upps! la contraseña no es válida")
		customError.SetStatusHTTP(http.StatusBadRequest)

		return model.User{}, customError
	}

	if strings.TrimSpace(m.FirstName) == "" || strings.TrimSpace(m.Lastname) == "" || strings.TrimSpace(m.Email) == "" {
		customError.SetCode(model.Failure)
		customError.SetAPIMessage("¡Upps! los datos no son válidos")
		customError.SetStatusHTTP(http.StatusBadRequest)

		return model.User{}, customError
	}

	passwordBcrypt, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		customError.SetCode(model.UnexpectedError)
		customError.SetAPIMessage("¡Upps! ocurrió un error inesperado")
		customError.SetStatusHTTP(http.StatusInternalServerError)
		customError.SetError(err)

		return model.User{}, fmt.Errorf("user: %w", customError)
	}
	m.Password = string(passwordBcrypt)

	err = u.storage.Create(m)
	if err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}
	m.Password = ""

	return m, nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	if email == "" {
		return model.User{}, fmt.Errorf("user: email is required")
	}

	user, err := u.storage.GetWhere(models.FieldsSpecification{
		Filters: models.Fields{{Name: "email", Value: email}},
	})
	if err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	return user, nil
}
