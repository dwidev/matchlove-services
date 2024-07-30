// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
	"matchlove-services/internal/handler"
	"matchlove-services/internal/repository"
	"matchlove-services/internal/router"
	"matchlove-services/internal/service"
	"matchlove-services/pkg/cache"
)

// Injectors from wire.go:

func InitializeHandler(db *gorm.DB, cache2 cache.Cache) *router.Handler {
	validate := validator.New()
	iAccountRepository := repository.NewAccountRepository(db)
	iUserRepository := repository.NewUserRepository(db, iAccountRepository)
	iUserService := service.NewUserService(iUserRepository)
	iUserHandler := handler.NewUserHandler(validate, iUserService)
	iAuthRepository := repository.NewAuthRepository(db, cache2)
	iAuthService := service.NewAuthService(iAuthRepository, iAccountRepository, iUserRepository)
	iAuthHandler := handler.NewAuthHandler(validate, iAuthService)
	iMatchmakingRepository := repository.NewMatchmakingRepository(db)
	iMatchmakingService := service.NewMatchMakingService(iMatchmakingRepository)
	iMatchmakingHandler := handler.NewMatchMakingHandler(validate, iMatchmakingService)
	routerHandler := &router.Handler{
		UserHandler:        iUserHandler,
		AuthHandler:        iAuthHandler,
		MatchmakingHandler: iMatchmakingHandler,
	}
	return routerHandler
}

// wire.go:

var (
	reportRepositorySet = wire.NewSet(repository.NewAccountRepository, repository.NewUserRepository, repository.NewAuthRepository, repository.NewMatchmakingRepository)

	serviceSet = wire.NewSet(service.NewUserService, service.NewAuthService, service.NewMatchMakingService)

	handlerSet = wire.NewSet(handler.NewAuthHandler, handler.NewUserHandler, handler.NewMatchMakingHandler)
)
