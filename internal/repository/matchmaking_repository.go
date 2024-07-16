package repository

import (
	"gorm.io/gorm"
	"matchlove-services/internal/model"
)

func NewMatchmakingRepository(db *gorm.DB) IMatchmakingRepository {
	return &MatchmakingRepository{
		db: db,
	}
}

type IMatchmakingRepository interface {
	GetMatchSuggestions(account string) ([]*model.UserAccount, error)
}

type MatchmakingRepository struct {
	db *gorm.DB
}

func (r *MatchmakingRepository) GetMatchSuggestions(account string) ([]*model.UserAccount, error) {
	result := make([]*model.UserAccount, 0)

	return result, nil
}
