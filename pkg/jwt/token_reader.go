package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"matchlove-services/internal/constant"
	"matchlove-services/pkg/middleware"
	"matchlove-services/pkg/response"
)

func getUuidFromToken(c *fiber.Ctx, contextKey string) (string, error) {
	user, ok := c.Locals(contextKey).(*jwt.Token)
	if !ok {
		return "", response.CredentialNoProvide
	}

	claim, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return "", response.TokenInvalidOrExpired
	}

	uuid, ok := claim[constant.UuidClaimKey].(string)
	if !ok {
		return "", response.TokenInvalidOrExpired
	}

	return uuid, nil
}

func GetUuidFromAccessToken(c *fiber.Ctx) (string, error) {
	return getUuidFromToken(c, middleware.ContextKeyAccess)
}

func GetUuidFromRefreshToken(c *fiber.Ctx) (string, error) {
	return getUuidFromToken(c, middleware.ContextKeyRefresh)
}
