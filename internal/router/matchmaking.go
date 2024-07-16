package router

import "matchlove-services/pkg/middleware"

func (r *Router) MatchMakingRoutes() {
	accessWare := middleware.JwtAccessProtected(r.Config)

	authRoute := r.v1.Group("/matches").Use(accessWare)
	authRoute.Get("/suggestions", r.Handler.MatchmakingHandler.GetMatchSuggestion)
}
