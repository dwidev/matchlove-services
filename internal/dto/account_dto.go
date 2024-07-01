package dto

type (
	AccountDTO struct {
		Uuid     string `json:"uuid"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	ChangePassswordDTO struct {
		OldPassword string `json:"oldPassword" validate:"required"`
		NewPassword string `json:"newPassword" validate:"required"`
	}
)
