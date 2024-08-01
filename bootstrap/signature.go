package bootstrap

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"

	"github.com/MikelSot/interseguro-challenge-auth/infrastructure/handler/request"
	"github.com/MikelSot/interseguro-challenge-auth/model"
)

func ValidateJWT(c *fiber.Ctx) error {
	customError := model.NewError()

	tokenHeader, err := request.GetTokenFromHeader(c)
	if err != nil {
		log.Warn("Se encontró un error al tratar de leer el token")

		customError.SetCode(model.Unauthorized)
		customError.SetAPIMessage(err.Error())
		customError.SetStatusHTTP(http.StatusUnauthorized)

		return fmt.Errorf("bootstrap: %w", customError)
	}

	verifyFunction := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Método de firma no válido")
		}

		return []byte(getSignKey()), nil
	}

	token, err := jwt.Parse(tokenHeader, verifyFunction)
	if errors.Is(err, jwt.ErrTokenExpired) {
		log.Warn("Token de acceso expirado")

		customError.SetCode(model.TokenExpired)
		customError.SetAPIMessage("Token de acceso expirado")
		customError.SetStatusHTTP(http.StatusUnauthorized)

		return customError
	}
	if err != nil {
		log.Warnf("Error al procesar el token %v", err)

		fmt.Println(getSignKey())

		customError.SetCode(model.UnexpectedError)
		customError.SetAPIMessage("Error al procesar el token")
		customError.SetStatusHTTP(http.StatusInternalServerError)

		return customError
	}
	if !token.Valid {
		log.Warn("Token de acceso no válido")

		customError.SetCode(model.Unauthorized)
		customError.SetAPIMessage("Token de acceso no válido")
		customError.SetStatusHTTP(http.StatusUnauthorized)

		return customError
	}

	return c.Next()
}
