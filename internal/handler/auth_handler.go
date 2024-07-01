package handler

import (
	"matchlove-services/internal/dto"
	"matchlove-services/internal/service"
	"matchlove-services/pkg/helper"
	"matchlove-services/pkg/jwt"
	"matchlove-services/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewAuthHandler(validate *validator.Validate, service service.IAuthService) IAuthHandler {
	return &authHandler{
		service:   service,
		Validator: validate,
	}
}

type IAuthHandler interface {
	LoginWithEmail(c *fiber.Ctx) error
	LoginWithEmailPassword(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
}

type authHandler struct {
	Validator *validator.Validate
	service   service.IAuthService
}

func (handler *authHandler) LoginWithEmail(c *fiber.Ctx) error {
	type Dto struct {
		Email string `json:"email" validate:"required,email"`
	}
	request := new(Dto)
	if err := c.BodyParser(request); err != nil {
		return response.CatchFiberError(err)
	}

	if listErr := helper.Validation(handler.Validator, request); len(listErr) != 0 {
		return response.FieldErrorResponse(c, listErr)
	}
	res, err := handler.service.OnLoginWithEmail(request.Email)

	if err != nil {
		return response.CatchFiberError(err)
	}

	return c.JSON(res)
}
func (handler *authHandler) LoginWithEmailPassword(c *fiber.Ctx) error {
	user := new(dto.LoginPassRequestDTO)
	if err := c.BodyParser(user); err != nil {
		return response.CatchFiberError(err)
	}

	if listErr := helper.Validation(handler.Validator, user); len(listErr) != 0 {
		return response.FieldErrorResponse(c, listErr)
	}
	res, err := handler.service.OnLoginWithEmailPassword(user)

	if err != nil {
		return response.CatchFiberError(err)
	}

	return c.JSON(res)
}

func (handler *authHandler) Logout(c *fiber.Ctx) error {
	userId, err := jwt.GetUuidFromAccessToken(c)
	if err != nil {
		return response.CatchFiberError(err)
	}

	err = handler.service.OnLogout(userId)
	if err != nil {
		return response.CatchFiberError(err)
	}

	return c.JSON(response.MessageResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Logout success",
	})
}

func (handler *authHandler) RefreshToken(c *fiber.Ctx) error {
	uuid, err := jwt.GetUuidFromRefreshToken(c)
	if err != nil {
		return response.CatchFiberError(err)
	}

	res, err := handler.service.RefreshToken(uuid)
	if err != nil {
		return response.CatchFiberError(err)
	}

	return c.JSON(res)
}

func (handler *authHandler) ChangePassword(c *fiber.Ctx) error {
	dto := new(dto.ChangePassswordDTO)
	if err := c.BodyParser(dto); err != nil {
		return response.CatchFiberError(err)
	}

	if listErr := helper.Validation(handler.Validator, dto); len(listErr) != 0 {
		return response.FieldErrorResponse(c, listErr)
	}

	id, err := jwt.GetUuidFromAccessToken(c)
	if err != nil {
		return response.CatchFiberError(err)
	}

	res, err := handler.service.ChangePassword(id, dto)
	if err != nil {
		return response.CatchFiberError(err)
	}

	return c.JSON(res)
}
