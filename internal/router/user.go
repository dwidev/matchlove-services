package router

import "matchlove-services/pkg/middleware"

func (r *Router) UserRoutes() {
	accessWare := middleware.JwtAccessProtected(r.Config)

	userRoute := r.v1.Group("/users").Use(accessWare)
	userRoute.Get("/me/profile", r.Handler.UserHandler.GetMyProfile)
	userRoute.Patch("/me/profile", r.Handler.UserHandler.UpdateProfile)
	userRoute.Get("/profile/:accountId", r.Handler.UserHandler.GetUserProfile)
}
