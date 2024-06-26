//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"matchlove-services/internal/handler"
	"matchlove-services/internal/repository"
	"matchlove-services/internal/router"
	"matchlove-services/internal/service"

	"github.com/google/wire"
)

var (
	reportRepositorySet = wire.NewSet(
		repository.NewAccountRepository,
		repository.NewUserRepository,
		repository.NewAuthRepository,
	)

	serviceSet = wire.NewSet(
		service.NewUserService,
		service.NewAuthService,
	)

	handlerSet = wire.NewSet(
		handler.NewAuthHandler,
		handler.NewUserHandler,
	)
)

func IntializeHandler(db *gorm.DB) *router.Handler {
	wire.Build(
		validator.New,
		reportRepositorySet,
		serviceSet,
		handlerSet,
		wire.Struct(new(router.Handler), "*"),
	)

	return nil
}
