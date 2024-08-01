package bootstrap

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/MikelSot/interseguro-challenge-auth/infrastructure/handler/request"
	"github.com/MikelSot/interseguro-challenge-auth/model"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

func LoadSignatures(private, public []byte, logger model.Logger) {
	once.Do(func() {
		var err error
		signKey, err = jwt.ParseRSAPrivateKeyFromPEM(private)
		if err != nil {
			logger.Fatalf("authorization.LoadSignatures: realizando el parse en jwt RSA private: %s", err)
		}

		verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(public)
		if err != nil {
			logger.Fatalf("authorization.LoadSignatures: realizando el parse en jwt RSA public: %s", err)
		}
	})
}

func ValidateJWT(c *fiber.Ctx) error {
	fmt.Println("Validating JWT")

	tokenHeader, err := request.GetTokenFromHeader(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Se encontró un error al tratar de leer el token")
	}

	verifyFunction := func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
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
