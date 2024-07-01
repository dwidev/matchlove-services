package dto

import "time"

type MerchantProfileDTO struct {
	Uuid         string      `json:"uuid"`
	MerchantName string      `json:"merchantName"`
	Description  string      `json:"description"`
	Location     string      `json:"location"`
	Account      *AccountDTO `json:"accountInfo,omitempty"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}
