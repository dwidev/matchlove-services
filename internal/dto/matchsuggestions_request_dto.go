package dto

type MatchSuggestionsRequestDto struct {
	AccountID string
	Page      int `query:"page"`
	PerPage   int `query:"per_page"`
}
