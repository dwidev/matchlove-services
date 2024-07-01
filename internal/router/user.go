package router

import (
	"matchlove-services/pkg/injection"
	"matchlove-services/pkg/middleware"

	"github.com/sirupsen/logrus"
)

func (r *Router) UserRoutes() {
	userHandler, err := injection.InitializeUserHandler(r.DB)
	if err != nil {
		logrus.Fatalf("error initializing user handler: %v", err)
	}

	accessWare := middleware.JwtAccessProtected(r.Config)
	userRoute := r.v1.Group("/user").Use(accessWare)
	userRoute.Post("/register", userHandler.RegisterUser)
}
