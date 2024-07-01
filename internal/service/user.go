package service

import (
	"matchlove-services/internal/dto"
	"matchlove-services/internal/repository"
)

func NewUserService(r repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: r,
	}
}

type IUserService interface {
	OnRegisterUser(req *dto.UserProfileRegisterDTO) error
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
