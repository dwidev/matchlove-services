package dto

import (
	"github.com/google/uuid"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/helper"
)

type UserProfileDTO struct {
	ID           string   `json:"id" validate:"required"`
	Bio          *string  `json:"bio" validate:"required"`
	Gender       *string  `json:"gender" validate:"required"`
	DateOfBirth  *string  `json:"date_of_birth" validate:"required"`
	Height       *float64 `json:"height" validate:"required"`
	Weight       *float64 `json:"weight" validate:"required"`
	Zodiac       *string  `json:"zodiac" validate:"required"`
	BloodType    *string  `json:"blood_type" validate:"required"`
	Education    *string  `json:"education" validate:"required"`
	Personality  *string  `json:"personality" validate:"required"`
	LookingFor   *string  `json:"looking_for" validate:"required"`
	LoveLanguage *string  `json:"love_language" validate:"required"`
	City         *string  `json:"city" validate:"required"`
}

func (u UserProfileDTO) ParseProfileToModel(accountID string) *model.UserProfile {
	return &model.UserProfile{
		Uuid:         uuid.MustParse(u.ID),
		AccountUuid:  accountID,
		Bio:          *u.Bio,
		Gender:       *u.Gender,
		DateOfBirth:  helper.ParseDateTime(*u.DateOfBirth),
		Height:       *u.Height,
		Weight:       *u.Weight,
		Zodiac:       *u.Zodiac,
		BloodType:    *u.BloodType,
		Education:    *u.Education,
		Personality:  *u.Personality,
		LookingFor:   *u.LookingFor,
		LoveLanguage: *u.LoveLanguage,
		City:         *u.City,
	}
}
