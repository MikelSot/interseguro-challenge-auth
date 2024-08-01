package register

import (
	"net/http"

	"github.com/MikelSot/interseguro-challenge-auth/model"
)

type Register struct {
	user  UserUseCase
	token TokenUseCase
}

func New(u UserUseCase, t TokenUseCase) Register {
	return Register{u, t}
}

func (r Register) Register(m model.User) (interface{}, error) {
	user, err := r.user.Create(m)
	if err != nil {
		return nil, err
	}

	customError := model.NewError()
	token, err := r.token.Generate(m)
	if err != nil {
		customError.SetCode(model.UnexpectedError)
		customError.SetAPIMessage("¡Upps! ocurrió un error inesperado")
		customError.SetStatusHTTP(http.StatusInternalServerError)

		return nil, customError
	}

	user.Password = ""

	mr := struct {
		model.User `json:"user"`
		Token      string `json:"token"`
	}{
		user,
		token,
	}

	return mr, nil
}
