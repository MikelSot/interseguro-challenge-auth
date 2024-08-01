package register

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/interseguro-challenge-auth/domain/register"
	"github.com/MikelSot/interseguro-challenge-auth/model"
)

type handler struct {
	useCase register.UseCase
}

func newHandler(uc register.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Register(c *fiber.Ctx) error {
	m := model.User{}

	if err := c.BodyParser(&m); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.MessageResponse{
			Errors: model.Responses{
				{Code: model.BindFailed, Message: "Error al leer el cuerpo de la solicitud"},
			},
		})
	}

	data, err := h.useCase.Register(m)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(data)
}
