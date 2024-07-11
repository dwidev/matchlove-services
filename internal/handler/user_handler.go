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
	GetMyProfile(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
	GetUserProfile(c *fiber.Ctx) error
}

type UserHandler struct {
	service   service.IUserService
	Validator *validator.Validate
}

func (handler *UserHandler) GetMyProfile(c *fiber.Ctx) error {
	accountID, err := jwt.GetUuidFromAccessToken(c)
	if err != nil {
		logrus.Errorf("UserHandler.GetProfile get account id from access token error: %v", err)
		return response.CatchFiberError(err)
	}

	account, err := handler.service.GetMyProfile(accountID)
	if err != nil {
		return response.CatchFiberError(err)
	}

	return response.SuccessResponse(c, account)
}

func (handler *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	request := new(dto.UpdateProfileRequestDTO)
	if err := c.BodyParser(request); err != nil {
		return response.CatchFiberError(err)
	}

	if listErr := helper.Validation(handler.Validator, request); len(listErr) != 0 {
		return response.FieldErrorResponse(c, listErr)
	}

	accountID, err := jwt.GetUuidFromAccessToken(c)
	if err != nil {
		logrus.Errorf("UserHandler.UpdateProfile get account id from access token error: %v", err)
		return response.CatchFiberError(err)
	}

	account, err := handler.service.UpdateProfile(accountID, request)
	if err != nil {
		return response.CatchFiberError(err)
	}

	return response.SuccessResponse(c, account)
}

func (handler *UserHandler) GetUserProfile(c *fiber.Ctx) error {
	accountID := c.Params("accountId")
	account, err := handler.service.GetUserProfile(accountID)
	if err != nil {
		return response.CatchFiberError(err)
	}

	return response.SuccessResponse(c, account)
}
