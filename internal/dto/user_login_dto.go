package dto

type LoginPassRequestDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SuccessLoginResponseDTO struct {
	StatusCode        int    `json:"statusCode"`
	AccessToken       string `json:"accessToken"`
	RefreshToken      string `json:"refreshToken"`
	IsNewAccount      bool   `json:"isNewAccount"`
	IsCompleteProfile bool   `json:"isCompleteProfile"`
}

type LoginWithEmailDto struct {
	Email string `json:"email" validate:"required,email"`
}
