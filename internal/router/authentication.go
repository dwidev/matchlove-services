package router

func (r *Router) AuthenticationRoutes() {
	// repository
	// accountRepository := repository.NewAccountRepository(r.DB)
	// userRepository := repository.NewUserRepository(r.DB, accountRepository)
	// authRepository := repository.NewAuthRepository(r.DB)

	// // service
	// authService := service.NewAuthService(authRepository, accountRepository)
	// // userService := service.NewUserService(userRepository)

	// // handler
	// // r.userHandler = handler.NewUserHandler(r.Validator, userService)
	// r.authHandler = handler.NewAuthHandler(r.Validator, authService)

	// refreshWare := middleware.JwtRefreshProtected(r.Config)
	// accessWare := middleware.JwtAccessProtected(r.Config)

	// authRoute := r.v1.Group("/auth")
	// authRoute.Post("/login-with-pass", r.authHandler.LoginWithEmailPassword)
	// authRoute.Post("/login-with-email", r.authHandler.LoginWithEmail)
	// authRoute.Post("/logout", accessWare, r.authHandler.Logout)
	// authRoute.Post("/change/password", accessWare, r.authHandler.ChangePassword)
	// authRoute.Post("/refresh", refreshWare, r.authHandler.RefreshToken)
}
