package request

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetTokenFromHeader(c *fiber.Ctx) (string, error) {
	auth := c.Get("Authorization")
	if auth == "" {
		return "", errors.New("el encabezado no contiene la autorización")
	}

	return auth, nil
}
