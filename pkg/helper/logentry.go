package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
)

func LogEntry(ctx *fiber.Ctx, moduleName *string) *logrus.Entry {
	defaultLog := logrus.Fields{
		"module": moduleName,
	}

	if ctx == nil {
		return logrus.WithFields(defaultLog)
	}

	var bodyData interface{}
	if err := ctx.BodyParser(&bodyData); err != nil {
		logrus.Warn("Failed to parse request body at MakeLogEntry: ", err)
		bodyData = "Failed to parse body"
	}

	defaultLog["module"] = "http_request"
	httpField := logrus.Fields{
		"status":     ctx.Response().StatusCode(),
		"uri":        ctx.OriginalURL(),
		"ip":         ctx.IP(),
		"user_agent": ctx.Get("User-Agent"),
		"data":       bodyData,
		"params":     string(ctx.Request().URI().QueryString()),
	}

	maps.Copy(defaultLog, httpField)
	return logrus.WithFields(defaultLog)
}
