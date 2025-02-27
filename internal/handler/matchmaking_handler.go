package handler

import (
	"fmt"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/internal/service"
	"matchlove-services/pkg/helper"
	"matchlove-services/pkg/jwt"
	"matchlove-services/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewMatchMakingHandler(validator *validator.Validate, service service.IMatchmakingService) IMatchmakingHandler {
	return &MatchmakingHandler{
		service:   service,
		validator: validator,
	}
}

type IMatchmakingHandler interface {
	GetMatchSuggestion(c *fiber.Ctx) error
	Like(c *fiber.Ctx) error
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

	request := new(dto.MatchSuggestionsRequestDto)
	if err = c.QueryParser(request); err != nil {
		return response.CatchFiberError(err)
	}

	if listErr := helper.Validation(handler.validator, request); len(listErr) > 0 {
		return response.FieldErrorResponse(c, listErr)
	}

	request.AccountID = accountID
	result, err := handler.service.GetMatchSuggestions(request)
	if err != nil {
		return response.CatchFiberError(err)
	}

	dummy := make([]string, 0)
	for _, account := range result.Data.([]*model.UserAccount) {
		dummy = append(dummy, fmt.Sprintf("%s#%s#%s", account.UserProfile.FirstName, account.Username, account.UserProfile.Uuid))
	}

	res := dto.PaginationResultDTO{
		CurrentPage: result.CurrentPage,
		TotalPage:   result.TotalPage,
		TotalData:   result.TotalData,
		Data:        result.Data,
	}

	return response.SuccessResponse(c, res)
}

func (handler *MatchmakingHandler) Like(c *fiber.Ctx) error {
	request := new(dto.LikeRequestDTO)
	if err := c.BodyParser(request); err != nil {
		return response.CatchFiberError(err)
	}

	if listErr := helper.Validation(handler.validator, request); len(listErr) > 0 {
		return response.FieldErrorResponse(c, listErr)
	}

	//uuid, err := jwt.GetUuidFromAccessToken(c)
	//if err != nil {
	//	return response.CatchFiberError(err)
	//}
	//
	//if uuid == request.SecondUserAccountID {
	//	return response.CatchFiberError(response.CannotLikeYourself)
	//}

	result, err := handler.service.Like(request)

	if err != nil {
		return response.CatchFiberError(err)
	}

	if result == dto.MATCHES {
		return response.SuccessResponse(c, "yeayyy it's a match")
	}

	return response.SuccessResponse(c, "user liked")
}
