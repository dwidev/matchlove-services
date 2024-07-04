package router

import "matchlove-services/pkg/middleware"

func (r *Router) UserRoutes() {
	accessWare := middleware.JwtAccessProtected(r.Config)

	userRoute := r.v1.Group("/users").Use(accessWare)
	userRoute.Post("/register", r.Handler.UserHandler.RegisterUser)
	userRoute.Get("/profile/:accountId", r.Handler.UserHandler.GetProfile)
}
