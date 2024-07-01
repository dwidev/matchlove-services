package jwt

import (
	"errors"
	"matchlove-services/internal/constant"
	"matchlove-services/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func getUuidFromToken(c *fiber.Ctx, contextKey string) (string, error) {
	user := c.Locals(contextKey).(*jwt.Token)

	claim, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	uuid, ok := claim[constant.UuidClaimKey].(string)
	if !ok {
		return "", errors.New("invalid uuid claims")
	}

	return uuid, nil
}

func GetUuidFromAccessToken(c *fiber.Ctx) (string, error) {
	return getUuidFromToken(c, middleware.ContextKeyAccess)
}

func GetUuidFromRefreshToken(c *fiber.Ctx) (string, error) {
	return getUuidFromToken(c, middleware.ContextKeyRefresh)
}
