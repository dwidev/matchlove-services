package service

import (
	"matchlove-services/internal/model"
	"matchlove-services/internal/repository"
)

func NewMatchMakingService(mmr repository.IMatchmakingRepository) IMatchmakingService {
	return &MatchMakingService{
		MatchmakingRepository: mmr,
	}
}

type IMatchmakingService interface {
	GetMatchSuggestions(accountId string) ([]*model.UserAccount, error)
}

type MatchMakingService struct {
	MatchmakingRepository repository.IMatchmakingRepository
}

func (m *MatchMakingService) GetMatchSuggestions(accountId string) ([]*model.UserAccount, error) {
	res, err := m.MatchmakingRepository.GetMatchSuggestions(accountId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
