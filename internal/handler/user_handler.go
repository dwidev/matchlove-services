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

func NewUserHandler(validator *validator.Validate, service service.IUserService) IUserHandler {
	return &UserHandler{
		service:   service,
		Validator: validator,
	}
}

type IUserHandler interface {
	RegisterUser(c *fiber.Ctx) error
}

type UserHandler struct {
	service   service.IUserService
	Validator *validator.Validate
}

func (handler *UserHandler) RegisterUser(c *fiber.Ctx) error {
	regisDto := new(dto.UserProfileRegisterDTO)
	if err := c.BodyParser(regisDto); err != nil {
		return response.ErrorResponse(c, err)
	}

	if listErr := helper.Validation(handler.Validator, regisDto); len(listErr) != 0 {
		return response.FieldErrorResponse(c, listErr)
	}

	accountid, err := jwt.GetUuidFromAccessToken(c)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	regisDto.AccountId = accountid
	if err := handler.service.OnRegisterUser(regisDto); err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, "Complete profile registered")
}
