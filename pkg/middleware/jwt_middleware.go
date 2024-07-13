package middleware

import (
	"github.com/sirupsen/logrus"
	"matchlove-services/internal/constant"
	"matchlove-services/pkg/config"
	"matchlove-services/pkg/response"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/gofiber/fiber/v2"
)

// JwtAccessProtected function for middleware with access token jwt
func JwtAccessProtected(env *config.Schema) fiber.Handler {
	c := jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(env.AccessSecretKey),
		},
		ContextKey: constant.ContextKeyAccess,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return jwtError(c, err, constant.ContextKeyAccess)
		},
	}

	return jwtware.New(c)
}

// JwtRefreshProtected function for middleware with refresh token jwt
func JwtRefreshProtected(env *config.Schema) fiber.Handler {
	c := jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(env.RefreshSecretKey),
		},
		ContextKey: constant.ContextKeyRefresh,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return jwtError(c, err, constant.ContextKeyRefresh)
		},
	}

	return jwtware.New(c)
}

// func for error handling jwt
func jwtError(c *fiber.Ctx, err error, types string) error {
	if strings.Contains(err.Error(), "missing") || strings.Contains(err.Error(), "malformed") {
		logrus.Errorf("error on jwt error handler with type %s, error : %s", types, err.Error())
		err := response.NewAppError(fiber.ErrBadRequest.Code, err.Error())
		return response.ErrorResponse(c, err)
	}

	err = response.NewAppError(fiber.StatusUnauthorized, err.Error())
	return response.ErrorResponse(c, err)
}
