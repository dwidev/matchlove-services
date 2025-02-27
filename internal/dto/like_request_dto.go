package dto

type LikeRequestDTO struct {
	FirstUserAccountID  string `json:"first_user_account_id" validate:"required"`
	SecondUserAccountID string `json:"second_user_account_id" validate:"required"`
}

type LikeResponseType string

const (
	ERROR   LikeResponseType = "error"
	LIKED   LikeResponseType = "liked"
	MATCHES LikeResponseType = "matches"
)
