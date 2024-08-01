package bootstrap

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/MikelSot/interseguro-challenge-auth/infrastructure/handler/request"
)

func ValidateJWT(c *fiber.Ctx) error {
	tokenHeader, err := request.GetTokenFromHeader(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Se encontró un error al tratar de leer el token")
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
		return fiber.NewError(fiber.StatusUnauthorized, "Token de acceso expirado")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error al procesar el token")
	}
	if !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Token de acceso no válido")
	}

	return c.Next()
}
