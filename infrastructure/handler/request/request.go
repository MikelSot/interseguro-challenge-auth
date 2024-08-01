package request

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetTokenFromHeader(c *fiber.Ctx) (string, error) {
	auth := c.Get("Authorization")
	if auth == "" {
		return "", errors.New("el encabezado no contiene la autorización")
	}

	if len(auth) > 6 && strings.ToUpper(auth[0:6]) == "BEARER" {
		return auth[7:], nil
	} else {
		return "", errors.New("el header no contiene la palabra Bearer")
	}
}
