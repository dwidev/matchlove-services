package middleware

import (
	"log"
	"matchlove-services/pkg/config"
	"matchlove-services/pkg/response"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/gofiber/fiber/v2"
)

const (
	ContextKeyAccess  = "JWT_ACCESS"
	ContextKeyRefresh = "JWT_REFRESH"
	ContextKeyUUID    = "USER_UUID"
)

// function for protect fobiddn access token
// func ForbiddenMiddleware(env *config.Schema, ac constant.AccountType) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		tokenFromHeader := c.Get(fiber.HeaderAuthorization)
// 		tokenFromHeader = tokenFromHeader[len("Bearer "):]

// 		token, err := jwt.Parse(tokenFromHeader, func(t *jwt.Token) (interface{}, error) {
// 			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected sigining method: %v", t.Header["alg"])
// 			}

// 			return []byte(env.AccessSecretKey), nil
// 		})

// 		if err != nil {
// 			log.Println("ERROR IN JWT PROTECTED", ac)
// 			log.Println(err)
// 			return response.ErrorResponse(c, &response.AppError{
// 				Code:    fiber.StatusUnauthorized,
// 				Message: "Invalid Token",
// 			})
// 		}
// 		claims := token.Claims

// 		if claims, ok := claims.(jwt.MapClaims); ok {
// 			accountType := claims[constant.AccountTypeClaimKey].(string)
// 			if constant.AccountType(accountType) != ac {
// 				return response.ErrorResponse(c, &response.AppError{
// 					Code:    fiber.StatusForbidden,
// 					Message: "Cannot access with credentials",
// 				})
// 			}

// 		}

// 		return c.Next()
// 	}
// }

// / function for middleware with access token jwt
func JwtAccessProtected(env *config.Schema) fiber.Handler {
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(env.AccessSecretKey),
		},
		ContextKey: ContextKeyAccess,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return jwtError(c, err, ContextKeyAccess)
		},
	}

	return jwtware.New(config)
}

// / function for middleware with refres token jwt
func JwtRefreshProtected(env *config.Schema) fiber.Handler {
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(env.RefreshSecretKey),
		},
		ContextKey: ContextKeyRefresh,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return jwtError(c, err, ContextKeyRefresh)
		},
	}

	return jwtware.New(config)
}

// func for error handling jwt
func jwtError(c *fiber.Ctx, err error, types string) error {
	log.Println("ERROR ON JWT ERROR HANDLER", types)
	log.Println(err)

	if strings.Contains(err.Error(), "missing or malformed JWT") {
		err := response.NewAppError(fiber.ErrBadRequest.Code, err.Error())
		return response.ErrorResponse(c, err)
	}

	err = response.NewAppError(fiber.StatusUnauthorized, err.Error())
	return response.ErrorResponse(c, err)

}
