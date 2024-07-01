package router

import "matchlove-services/pkg/middleware"

func (r *Router) UserRoutes() {
	accessWare := middleware.JwtAccessProtected(r.Config)

	userRoute := r.v1.Group("/user").Use(accessWare)
	userRoute.Post("/register", r.Handler.UserHandler.RegisterUser)
}
