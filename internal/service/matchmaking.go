package service

import (
	"matchlove-services/internal/dto"
	"matchlove-services/internal/repository"
)

func NewMatchMakingService(mmr repository.IMatchmakingRepository) IMatchmakingService {
	return &MatchMakingService{
		MatchmakingRepository: mmr,
	}
}

type IMatchmakingService interface {
	GetMatchSuggestions(request *dto.MatchSuggestionsRequestDto) (*dto.PaginationResultDTO, error)
	Like(request *dto.LikeRequestDTO) (dto.LikeResponseType, error)
}

type MatchMakingService struct {
	MatchmakingRepository repository.IMatchmakingRepository
}

func (m *MatchMakingService) GetMatchSuggestions(request *dto.MatchSuggestionsRequestDto) (*dto.PaginationResultDTO, error) {
	res, err := m.MatchmakingRepository.GetMatchSuggestions(request)
	//res = helper.RandomArray(res, len(res))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MatchMakingService) Like(request *dto.LikeRequestDTO) (dto.LikeResponseType, error) {
	// check second user already liked or no
	// if yes, it's a match
	// if no create like

	// create like
	_, err := m.MatchmakingRepository.CreateLike(request)
	if err != nil {
		return dto.ERROR, err
	}

	/// check match
	var matches, errMatch = m.MatchmakingRepository.CheckForMatches(request)
	if errMatch != nil {
		return dto.ERROR, err
	}

	/// create matches if match
	if matches == true {
		_, err := m.MatchmakingRepository.CreateMatches(request)
		if err != nil {
			return dto.ERROR, err
		}

		return dto.MATCHES, nil
	}

	return dto.LIKED, nil
}
