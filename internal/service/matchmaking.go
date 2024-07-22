package service

import (
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/internal/repository"
)

func NewMatchMakingService(mmr repository.IMatchmakingRepository) IMatchmakingService {
	return &MatchMakingService{
		MatchmakingRepository: mmr,
	}
}

type IMatchmakingService interface {
	GetMatchSuggestions(request *dto.MatchSuggestionsRequestDto) ([]*model.UserAccount, error)
}

type MatchMakingService struct {
	MatchmakingRepository repository.IMatchmakingRepository
}

func (m *MatchMakingService) GetMatchSuggestions(request *dto.MatchSuggestionsRequestDto) ([]*model.UserAccount, error) {
	res, err := m.MatchmakingRepository.GetMatchSuggestions(request)
	//res = helper.RandomArray(res, len(res))

	if err != nil {
		return nil, err
	}

	return res, nil
}
