package response

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type DataResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

type MessageResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
}

func NewResponse(statusCode int, Data interface{}) *DataResponse {
	return &DataResponse{
		StatusCode: statusCode,
		Data:       Data,
	}
}

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	response := NewResponse(fiber.StatusOK, data)
	return c.JSON(response)
}

func ErrorResponse(c *fiber.Ctx, err error) error {
	var appErr *AppError
	if errors.As(err, &appErr) && appErr.Code != fiber.StatusInternalServerError {
		response := &MessageResponse{
			StatusCode: appErr.Code,
			Message:    appErr.Message,
		}
		return c.Status(appErr.Code).JSON(response)
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) && fiberErr.Code != fiber.StatusInternalServerError {
		response := &MessageResponse{
			StatusCode: fiberErr.Code,
			Message:    fiberErr.Message,
		}
		return c.Status(fiberErr.Code).JSON(response)
	}

	response := &MessageResponse{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "There was an error on the server, please try again later.",
	}
	return c.Status(response.StatusCode).JSON(response)
}

func FieldErrorResponse(c *fiber.Ctx, listErrField []string) error {
	err := &AppError{
		Code:    fiber.StatusBadRequest,
		Message: listErrField,
	}
	return ErrorResponse(c, err)
}
