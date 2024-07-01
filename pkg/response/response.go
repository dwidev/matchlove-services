package response

import (
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

func NewErrorResponse(err error) *MessageResponse {
	appErr, ok := err.(*AppError)
	if ok {
		return &MessageResponse{
			StatusCode: appErr.Code,
			Message:    appErr.Message,
		}
	}

	return &MessageResponse{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "There was an error on the server, please try again later.",
	}
}

func ErrorResponse(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*AppError); ok {
		response := NewErrorResponse(appErr)
		return c.Status(appErr.Code).JSON(response)
	}

	response := NewErrorResponse(err)
	return c.Status(fiber.StatusInternalServerError).JSON(response)
}

func FieldErrorResponse(c *fiber.Ctx, listErrField []string) error {
	err := &AppError{
		Code:    fiber.StatusBadRequest,
		Message: listErrField,
	}
	return ErrorResponse(c, err)
}
