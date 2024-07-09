package dto

import (
	"matchlove-services/internal/model"
)

type UserRoutineDTO struct {
	MyWeekend      *string `json:"my_weekend" validate:"required"`
	MyHangouts     *string `json:"my_hangouts" validate:"required"`
	MorningRoutine *string `json:"morning_routine" validate:"required"`
}

func (r *UserRoutineDTO) ParseToModel(accountID string) *model.UserRoutine {
	return &model.UserRoutine{
		AccountID:      accountID,
		MyWeekend:      *r.MyWeekend,
		MyHangouts:     *r.MyHangouts,
		MorningRoutine: *r.MorningRoutine,
	}
}
