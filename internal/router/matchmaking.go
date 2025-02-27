package router

func (r *Router) MatchMakingRoutes() {
	authRoute := r.v1.Group("/matches")
	authRoute.Get("/suggestions", r.Handler.MatchmakingHandler.GetMatchSuggestion)
	authRoute.Get("/like", r.Handler.MatchmakingHandler.Like)
}
