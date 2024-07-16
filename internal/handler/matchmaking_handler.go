package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"matchlove-services/internal/service"
	"matchlove-services/pkg/jwt"
	"matchlove-services/pkg/response"
)

func NewMatchMakingHandler(validator *validator.Validate, service service.IMatchmakingService) IMatchmakingHandler {
	return &MatchmakingHandler{
		service:   service,
		validator: validator,
	}
}

type IMatchmakingHandler interface {
	GetMatchSuggestion(c *fiber.Ctx) error
}

type MatchmakingHandler struct {
	validator *validator.Validate
	service   service.IMatchmakingService
}

func (handler *MatchmakingHandler) GetMatchSuggestion(c *fiber.Ctx) error {
	accountID, err := jwt.GetUuidFromAccessToken(c)
	if err != nil {
		return response.CatchFiberError(err)
	}

	result, err := handler.service.GetMatchSuggestions(accountID)
	if err != nil {
		return response.CatchFiberError(err)
	}
	return response.SuccessResponse(c, result)
}
