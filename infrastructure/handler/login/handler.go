package login

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/MikelSot/interseguro-challenge-auth/domain/login"
	"github.com/MikelSot/interseguro-challenge-auth/model"
)

type handler struct {
	useCase login.UseCase
}

func newHandler(uc login.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Login(c *fiber.Ctx) error {
	m := model.User{}

	if err := c.BodyParser(&m); err != nil {
		log.Warn("¡Uy! Error al leer el cuerpo de la solicitud", err.Error())

		return c.Status(fiber.StatusBadRequest).JSON(model.MessageResponse{
			Errors: model.Responses{
				{Code: model.BindFailed, Message: "Error al leer el cuerpo de la solicitud"},
			},
		})
	}

	data, err := h.useCase.Login(m)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(data)
}
