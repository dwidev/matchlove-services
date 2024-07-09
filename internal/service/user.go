package service

import (
	"github.com/google/uuid"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/internal/repository"
)

func NewUserService(r repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: r,
	}
}

type IUserService interface {
	GetProfile(accountID string) (*model.UserAccount, error)
	UpdateProfile(accountID string, requestDTO *dto.UpdateProfileRequestDTO) (*model.UserAccount, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func (u *UserService) GetProfile(accountID string) (*model.UserAccount, error) {
	account, err := u.userRepository.GetProfile(accountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (u *UserService) UpdateProfile(accountID string, requestDTO *dto.UpdateProfileRequestDTO) (*model.UserAccount, error) {
	account := new(model.UserAccount)
	account.Uuid = uuid.MustParse(accountID)
	account.UserProfile = requestDTO.UserProfile.ParseProfileToModel(accountID)
	for _, interest := range requestDTO.UserInterest {
		account.UserInterest = append(account.UserInterest, interest.ParseToModel(accountID))
	}
	account.UserLifeStyle = requestDTO.UserLifestyle.ParseToModel(accountID)
	account.UserRoutine = requestDTO.UserRoutine.ParseToModel(accountID)

	account, err := u.userRepository.UpdateProfile(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
