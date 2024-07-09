package dto

import "matchlove-services/internal/model"

type UserInterestDto struct {
	Code string `json:"code"`
}

func (i UserInterestDto) ParseToModel(accountID string) *model.UserInterest {
	return &model.UserInterest{
		AccountID:    accountID,
		InterestCode: i.Code,
	}
}
