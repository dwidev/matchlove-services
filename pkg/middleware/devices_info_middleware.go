package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"matchlove-services/internal/constant"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/response"
)

func DeviceInfo(c *fiber.Ctx) error {
	listErr := make([]string, 0)
	imei := c.Get("imei")
	platform := c.Get("platform")
	osVersion := c.Get("os-version")
	deviceName := c.Get("device-name")

	if imei == "" {
		listErr = append(listErr, "imei is required")
	} else if platform == "" {
		listErr = append(listErr, "platform is required")
	} else if osVersion == "" {
		listErr = append(listErr, "os-version is required")
	} else if deviceName == "" {
		listErr = append(listErr, "device name is required")
	}

	if len(listErr) > 0 {
		return response.FieldErrorResponse(c, listErr)
	}

	deviceInfo := &model.DevicesInfo{
		ID:         uuid.New(),
		Imei:       imei,
		Platform:   platform,
		OsVersion:  osVersion,
		DeviceName: deviceName,
	}

	c.Locals(constant.DeviceInfoContextKey, deviceInfo)
	return c.Next()
}

func GetDeviceInfo(ctx *fiber.Ctx) *model.DevicesInfo {
	deviceInfo, ok := ctx.Locals(constant.DeviceInfoContextKey).(*model.DevicesInfo)
	if !ok {
		panic("(GetDeviceInfo) failed to get device info")
	}

	return deviceInfo
}
