package handler

import (
	"github.com/sirupsen/logrus"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/internal/service"
	"matchlove-services/pkg/helper"
	"matchlove-services/pkg/jwt"
	"matchlove-services/pkg/middleware"
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
	RegisterUser(c *fiber.Ctx) error
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
	request := new(dto.LoginWithEmailDto)
	if err := c.BodyParser(request); err != nil {
		return response.CatchFiberError(err)
	}

	if listErr := helper.Validation(handler.Validator, request); len(listErr) != 0 {
		return response.FieldErrorResponse(c, listErr)
	}

	deviceInfo := middleware.GetDeviceInfo(c)
	request.RecordLogin = &dto.RecordLoginActivityDto{
		LoginActivity: model.NewLoginActivity(),
		DevicesInfo:   deviceInfo,
	}
	res, err := handler.service.OnLoginWithEmail(request)
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

func (handler *authHandler) RegisterUser(c *fiber.Ctx) error {
	regisDto := new(dto.UserProfileRegisterDTO)
	if err := c.BodyParser(regisDto); err != nil {
		return response.ErrorResponse(c, err)
	}

	if listErr := helper.Validation(handler.Validator, regisDto); len(listErr) != 0 {
		return response.FieldErrorResponse(c, listErr)
	}

	accountID, err := jwt.GetUuidFromAccessToken(c)
	if err != nil {
		logrus.Errorf("authHandler.RegisterUser get account id from access token error: %v", err)
		return response.CatchFiberError(err)
	}

	regisDto.AccountId = accountID
	if err := handler.service.OnRegisterUser(regisDto); err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, "Complete profile registered")
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
	accountID, err := jwt.GetUuidFromRefreshToken(c)
	if err != nil {
		return response.CatchFiberError(err)
	}

	refreshToken := jwt.GetRefreshToken(c)
	res, err := handler.service.RefreshToken(accountID, refreshToken)
	if err != nil {
		return response.CatchFiberError(err)
	}

	return c.JSON(res)
}

func (handler *authHandler) ChangePassword(c *fiber.Ctx) error {
	request := new(dto.ChangePassswordDTO)
	if err := c.BodyParser(request); err != nil {
		return response.CatchFiberError(err)
	}

	if listErr := helper.Validation(handler.Validator, request); len(listErr) != 0 {
		return response.FieldErrorResponse(c, listErr)
	}

	id, err := jwt.GetUuidFromAccessToken(c)
	if err != nil {
		return response.CatchFiberError(err)
	}

	res, err := handler.service.ChangePassword(id, request)
	if err != nil {
		return response.CatchFiberError(err)
	}

	return c.JSON(res)
}
