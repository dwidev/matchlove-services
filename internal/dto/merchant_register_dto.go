package dto

type MerchantRegisterDTO struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Email        string `json:"email"`
	MerchantName string `json:"merchantName" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Location     string `json:"location" validate:"required"`
}
