package login

import (
	"database/sql"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/MikelSot/interseguro-challenge-auth/model"
)

type Login struct {
	user  UserUseCase
	token TokenUseCase
}

func New(user UserUseCase, token TokenUseCase) Login {
	return Login{user, token}
}

func (l Login) Login(m model.User) (interface{}, error) {
	customError := model.NewError()

	user, err := l.user.GetByEmail(m.Email)
	if errors.Is(err, sql.ErrNoRows) {
		customError.SetCode(model.RecordNotFound)
		customError.SetAPIMessage("¡Upps! el usuario no existe")
		customError.SetStatusHTTP(http.StatusNotFound)
		customError.SetError(err)

		return nil, customError
	}
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m.Password))
	if err != nil {
		customError.SetCode(model.InvalidPassword)
		customError.SetAPIMessage("¡Upps! la contraseña no es correcta")
		customError.SetStatusHTTP(http.StatusUnauthorized)

		return nil, customError
	}
	m.Password = ""
	user.Password = ""

	token, err := l.token.Generate(m)
	if err != nil {
		customError.SetCode(model.UnexpectedError)
		customError.SetAPIMessage("¡Upps! ocurrió un error inesperado")
		customError.SetStatusHTTP(http.StatusInternalServerError)

		return nil, customError
	}

	mr := struct {
		model.User `json:"user"`
		Token      string `json:"token"`
	}{
		user,
		token,
	}

	return mr, nil
}
