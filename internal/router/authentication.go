package router

import "matchlove-services/pkg/middleware"

func (r *Router) AuthenticationRoutes() {
	refreshWare := middleware.JwtRefreshProtected(r.Config)
	accessWare := middleware.JwtAccessProtected(r.Config)

	authRoute := r.v1.Group("/auth")

	authRoute.Post("/login-with-pass", r.Handler.AuthHandler.LoginWithEmailPassword)
	authRoute.Post("/login-with-email", middleware.DeviceInfo, r.Handler.AuthHandler.LoginWithEmail)
	authRoute.Post("/register", accessWare, r.Handler.AuthHandler.RegisterUser)
	authRoute.Post("/logout", accessWare, middleware.DeviceInfo, r.Handler.AuthHandler.Logout)
	authRoute.Post("/change/password", accessWare, r.Handler.AuthHandler.ChangePassword)
	authRoute.Post("/refresh", refreshWare, r.Handler.AuthHandler.RefreshToken)
}
