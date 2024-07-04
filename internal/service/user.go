package service

import (
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
	OnRegisterUser(req *dto.UserProfileRegisterDTO) error
	GetProfile(accountID string) (*model.UserAccount, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func (u *UserService) OnRegisterUser(req *dto.UserProfileRegisterDTO) error {
	err := u.userRepository.RegisterUser(req)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetProfile(accountID string) (*model.UserAccount, error) {
	account, err := u.userRepository.GetProfile(accountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}
