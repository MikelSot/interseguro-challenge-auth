package token

import (
	"crypto/rsa"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/MikelSot/interseguro-challenge-auth/model"
)

type Token struct {
	ExpiresAt int

	_signKey *rsa.PrivateKey
}

func New(expiresAt int, signKey *rsa.PrivateKey) Token {
	return Token{expiresAt, signKey}
}

func (t Token) Generate(m model.User) (string, error) {
	payload := jwt.MapClaims{
		"id":    m.ID,
		"email": m.Email,
		"ia":    time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * time.Duration(t.ExpiresAt)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	key, err := token.SignedString(t._signKey)
	if err != nil {
		customError := model.NewError()

		customError.SetCode(model.UnexpectedError)
		customError.SetAPIMessage("¡Upps! ocurrió un error inesperado")
		customError.SetStatusHTTP(http.StatusInternalServerError)

		return "", customError
	}

	return key, nil
}
