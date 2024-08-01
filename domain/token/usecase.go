package token

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/MikelSot/interseguro-challenge-auth/model"
)

type Token struct {
	ExpiresAt int

	SignKey string
}

func New(expiresAt int, signKey string) Token {
	return Token{expiresAt, signKey}
}

func (t Token) Generate(m model.User) (string, error) {
	payload := jwt.MapClaims{
		"email": m.Email,
		"ia":    time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * time.Duration(t.ExpiresAt)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	key, err := token.SignedString([]byte(t.SignKey))
	if err != nil {
		customError := model.NewError()

		customError.SetCode(model.UnexpectedError)
		customError.SetAPIMessage("¡Upps! ocurrió un error inesperado")
		customError.SetStatusHTTP(http.StatusInternalServerError)

		return "", customError
	}

	return key, nil
}
