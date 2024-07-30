package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"matchlove-services/pkg/cache"
	"matchlove-services/pkg/jwt"
	"matchlove-services/pkg/response"
)

func CheckToken(caching cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		deviceInfo := GetDeviceInfo(c)
		accountID, err := jwt.GetUuidFromAccessToken(c)
		if err != nil {
			return response.CatchFiberError(err)
		}

		key := cache.AccessTokenKeyCache(accountID, deviceInfo.Imei)
		ctx := c.Context()
		value, err := caching.GetString(ctx, key)

		if errors.Is(err, redis.Nil) {
			return response.CatchFiberError(response.CredentialNoLongerValid)
		}

		if err != nil {
			return response.CatchFiberError(err)
		}

		if len(value) < 0 {
			return response.ErrorResponse(c, response.CredentialNoProvide)
		}

		return c.Next()
	}
}
