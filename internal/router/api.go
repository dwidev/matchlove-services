package router

import (
	"matchlove-services/pkg/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Router struct {
	Validator *validator.Validate
	Engine    *fiber.App
	Config    *config.Schema
	DB        *gorm.DB

	v1 fiber.Router
}

func (r *Router) Build() {
	api := r.Engine.Group("/matchlove/api")
	r.v1 = api.Group("/v1/")

	r.AuthenticationRoutes()
	r.UserRoutes()
}
