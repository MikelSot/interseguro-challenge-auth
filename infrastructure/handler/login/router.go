package login

import (
	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/interseguro-challenge-auth/domain/login"
	"github.com/MikelSot/interseguro-challenge-auth/domain/token"
	"github.com/MikelSot/interseguro-challenge-auth/domain/user"
	userStorage "github.com/MikelSot/interseguro-challenge-auth/infrastructure/postgresql/user"
	"github.com/MikelSot/interseguro-challenge-auth/model"
)

const (
	_publicRoutePrefix = "/auth/api/v1/login"
)

func NewRouter(spec model.RouterSpecification) {
	handler := buildHandler(spec)

	publicRoutes(spec.App, handler)

}

func buildHandler(spec model.RouterSpecification) handler {

	userUseCase := user.New(userStorage.New(spec.DB))
	tokenUseCase := token.New(spec.ExpiresAt, spec.SignKey)

	useCase := login.New(userUseCase, tokenUseCase)

	return newHandler(useCase)
}
func publicRoutes(app *fiber.App, handler handler, middlewares ...fiber.Handler) {
	api := app.Group(_publicRoutePrefix, middlewares...)

	api.Post("", handler.Login)
}
