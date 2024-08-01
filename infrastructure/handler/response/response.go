package response

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/interseguro-challenge-auth/model"
)

// ErrorHandler handler the error response of fiber
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// custom error
	var e *model.Error
	ok := errors.As(err, &e)
	if ok {
		status, message := getErrorResponse(e)

		return ctx.Status(status).JSON(message)
	}

	// if the handler not returns a "model.Error" then it returns a generic error JSON response
	return ctx.Status(http.StatusInternalServerError).JSON(model.MessageResponse{
		Errors: model.Responses{
			{Code: model.UnexpectedError, Message: "¡Uy! metimos la pata, disculpanos lo solucionaremos"},
		},
	})
}

// getResponseError returns the status code and a Response
func getErrorResponse(err *model.Error) (int, model.MessageResponse) {
	if !err.HasCode() {
		err.SetCode(model.UnexpectedError)
	}

	if !err.HasAPIMessage() {
		err.SetErrorAsAPIMessage()
	}

	outputResponse := model.MessageResponse{}
	if err.HasData() {
		outputResponse.Data = err.Data()
	}

	if !err.HasStatusHTTP() {
		err.SetStatusHTTP(http.StatusInternalServerError)
	}

	outputStatus := err.StatusHTTP()
	outputResponse.Errors = model.Responses{model.Response{
		Code:    err.Code(),
		Message: err.APIMessage(),
	}}

	return outputStatus, outputResponse
}
