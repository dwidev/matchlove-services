package dto

import (
	"strings"
)

type UserProfileRegisterDTO struct {
	AccountId   string
	Name        string   `json:"name" validate:"required"`
	Gender      string   `json:"gender" validate:"required"`
	DOB         string   `json:"dateOfBirth" validate:"required"`
	Age         uint8    `json:"age" validate:"required"`
	Distance    int8     `json:"distance" validate:"required"`
	LookingFor  string   `json:"lookingFor"`
	InterestFor []string `json:"interestFor"`
	Longitude   float64  `json:"longitude"`
	Latitude    float64  `json:"latitude"`
}

func (u UserProfileRegisterDTO) JoinInterest() string {
	res := strings.Join(u.InterestFor, "#")
	return res
}

func (u UserProfileRegisterDTO) ToPreferredGender() string {
	if u.Gender == "Male" {
		return "Female"
	} else {
		return "Male"
	}
}
