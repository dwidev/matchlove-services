package dto

type MerchantUpdateDTO struct {
	MerchantName string `json:"merchantName"`
	Description  string `json:"description"`
	Location     string `json:"location"`
}
