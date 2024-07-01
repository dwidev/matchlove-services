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
}

func (u UserProfileRegisterDTO) JoinInterest() string {
	res := strings.Join(u.InterestFor, "#")
	return res
}

func (u UserProfileRegisterDTO) ToPreferedGender() string {
	if u.Gender == "Male" {
		return "Female"
	} else {
		return "Male"
	}
}
