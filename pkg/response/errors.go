package response

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func (e AppError) Error() string {
	if message, ok := e.Message.(string); ok {
		return message
	}

	if slice, ok := e.Message.([]string); ok {
		message := strings.Join(slice, ", ")
		return message
	}

	return fmt.Sprintf("%v", e.Message)
}

var (
	ErrUnauthorized  = &AppError{Code: fiber.StatusUnauthorized, Message: "Unauthorized"}
	PassNoValid      = &AppError{Code: fiber.StatusBadRequest, Message: "This password is wrong"}
	OldPasswordWrong = &AppError{Code: fiber.StatusBadRequest, Message: "Old password is wrong"}
	AccountNotFound  = &AppError{Code: fiber.StatusBadRequest, Message: "No active account found with the given credentials"}
	AlreadyExist     = &AppError{Code: fiber.StatusFound, Message: "Account is registered"}
	RecordNotFound   = &AppError{Code: fiber.StatusNotFound, Message: "Data is not found"}
)

func BadRequest(message interface{}) *AppError {
	return NewAppError(fiber.StatusBadRequest, message)
}

func NewAppError(code int, message interface{}) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func CatchFiberError(err error) error {
	var appErr *AppError
	if errors.As(err, &appErr) && appErr.Code != fiber.StatusInternalServerError {
		return appErr
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) && fiberErr.Code != fiber.StatusInternalServerError {
		return fiber.NewError(fiberErr.Code, fiberErr.Error())
	}

	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
}

func ValidateError(errorMsg []string) error {
	return fiber.ErrBadRequest
}
