package router

import (
	"github.com/gofiber/fiber/v2"
	"matchlove-services/internal/handler"
	"matchlove-services/pkg/config"
)

type Handler struct {
	UserHandler handler.IUserHandler
	AuthHandler handler.IAuthHandler
}

type Router struct {
	Config  *config.Schema
	Engine  *fiber.App
	Handler *Handler

	v1                      fiber.Router
	accessWare, refreshWare *fiber.Handler
}

func (r *Router) Build() {
	api := r.Engine.Group("/matchlove/api")
	r.v1 = api.Group("/v1/")

	r.AuthenticationRoutes()
	r.UserRoutes()
}
