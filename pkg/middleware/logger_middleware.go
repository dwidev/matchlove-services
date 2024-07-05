package middleware

import (
	"matchlove-services/pkg/helper"
	"matchlove-services/pkg/response"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

func Logging(ctx *fiber.Ctx) error {
	if err := ctx.Next(); err != nil {
		helper.LogEntry(ctx, nil).Errorf("ERROR:%s", err)
		return response.ErrorResponse(ctx, err)
	}

	helper.LogEntry(ctx, nil).Info("Incoming http request")
	return nil
}

func RecoverPanicLogging() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			logrus.Errorf("Panic: %v\nStack: %s", e, debug.Stack())
		},
	})
}
