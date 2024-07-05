package handler

import (
	"github.com/sirupsen/logrus"
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
	GetProfile(c *fiber.Ctx) error
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

	accountID, err := jwt.GetUuidFromAccessToken(c)
	if err != nil {
		logrus.Errorf("UserHandler.RegisterUser get account id from access token error: %v", err)
		return response.CatchFiberError(err)
	}

	regisDto.AccountId = accountID
	if err := handler.service.OnRegisterUser(regisDto); err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, "Complete profile registered")
}

func (handler *UserHandler) GetProfile(c *fiber.Ctx) error {
	accountID, err := jwt.GetUuidFromAccessToken(c)
	if err != nil {
		logrus.Errorf("UserHandler.GetProfile get account id from access token error: %v", err)
		return response.CatchFiberError(err)
	}

	account, err := handler.service.GetProfile(accountID)
	if err != nil {
		return response.CatchFiberError(err)
	}

	return response.SuccessResponse(c, account)
}
