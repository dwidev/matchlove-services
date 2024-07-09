package dto

import (
	"matchlove-services/internal/model"
)

type UserLifestyleDTO struct {
	SocialMediaActivity *string `json:"social_media_activity" validate:"required"`
	Pets                *string `json:"pets" validate:"required"`
	Drinking            *string `json:"drinking" validate:"required"`
	Smoking             *string `json:"smoking" validate:"required"`
	Workout             *string `json:"workout" validate:"required"`
	SleepingHabits      *string `json:"sleeping_habits" validate:"required"`
}

func (i *UserLifestyleDTO) ParseToModel(accountID string) *model.UserLifeStyle {
	ls := &model.UserLifeStyle{
		AccountID:           accountID,
		SocialMediaActivity: *i.SocialMediaActivity,
		Pets:                *i.Pets,
		Drinking:            *i.Drinking,
		Smoking:             *i.Smoking,
		Workout:             *i.Workout,
		SleepingHabits:      *i.SleepingHabits,
	}

	return ls
}
