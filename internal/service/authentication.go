package service

import (
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/internal/repository"
	"matchlove-services/pkg/helper"
	"matchlove-services/pkg/jwt"
	"matchlove-services/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewAuthService(
	r repository.AuthRepository,
	ar repository.AccountRepository,
) IAuthService {
	return &AuthService{
		AuthRepository:    r,
		AccountRepository: ar,
	}
}

type IAuthService interface {
	OnLoginWithEmail(email string) (*dto.SuccessLoginResponseDTO, error)
	OnLoginWithEmailPassword(req *dto.LoginPassRequestDTO) (*dto.SuccessLoginResponseDTO, error)
	OnLogout(userId string) error
	RefreshToken(userID string) (*dto.SuccessLoginResponseDTO, error)
	ChangePassword(userID string, dto *dto.ChangePassswordDTO) (*response.MessageResponse, error)
}

type AuthService struct {
	repository.AuthRepository
	repository.AccountRepository
}

func (s *AuthService) OnLoginWithEmail(email string) (*dto.SuccessLoginResponseDTO, error) {
	newAccount := false
	account, err := s.AuthRepository.LoginWithEmail(email)
	if err != nil {
		return nil, err
	}

	// handle user never login
	if account == nil {
		// for login with email if account not found set username with email
		accountRequest := &model.UserAccount{
			Uuid:     uuid.New(),
			Email:    email,
			Username: email,
		}
		account, err = s.AuthRepository.CreateNewAccount(accountRequest)
		newAccount = true
		if err != nil {
			return nil, err
		}
	}

	token, err := jwt.GenerateToken(account)
	if err != nil {
		return nil, err
	}

	err = s.AccountRepository.UpdateAccountRefreshToken(account.Uuid.String(), token.RefreshToken)
	if err != nil {
		return nil, err
	}

	result := &dto.SuccessLoginResponseDTO{
		StatusCode:        fiber.StatusOK,
		AccessToken:       token.AccessToken,
		RefreshToken:      token.RefreshToken,
		IsNewAccount:      newAccount,
		IsCompleteProfile: account.IsCompleteProfile == 1,
	}

	return result, nil
}

func (s *AuthService) OnLoginWithEmailPassword(req *dto.LoginPassRequestDTO) (*dto.SuccessLoginResponseDTO, error) {
	account, err := s.AuthRepository.LoginWithPassword(req)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(account)
	if err != nil {
		return nil, err
	}

	err = s.AccountRepository.UpdateAccountRefreshToken(account.Uuid.String(), token.RefreshToken)
	if err != nil {
		return nil, err
	}

	result := &dto.SuccessLoginResponseDTO{
		StatusCode:   fiber.StatusOK,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return result, nil
}

func (s *AuthService) OnLogout(userId string) error {
	err := s.AuthRepository.Logout(userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) RefreshToken(userID string) (*dto.SuccessLoginResponseDTO, error) {
	account, err := s.AccountRepository.GetAccountByID(userID)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(account)

	if err != nil {
		return nil, err
	}

	err = s.AccountRepository.UpdateAccountRefreshToken(userID, token.RefreshToken)
	if err != nil {
		return nil, err
	}

	result := &dto.SuccessLoginResponseDTO{
		StatusCode:   fiber.StatusOK,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return result, nil
}

func (s *AuthService) ChangePassword(userID string, dto *dto.ChangePassswordDTO) (*response.MessageResponse, error) {
	account, err := s.AccountRepository.GetAccountByID(userID)
	if err != nil {
		return nil, err
	}

	_, err = helper.ToCompare(dto.OldPassword, account.Password)
	if err != nil {
		return nil, response.OldPasswordWrong
	}

	newPassword, err := helper.ToHash(dto.NewPassword)
	if err != nil {
		return nil, err
	}

	err = s.AccountRepository.ChangePassword(userID, newPassword)
	if err != nil {
		return nil, err
	}

	return &response.MessageResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Password has changes",
	}, nil
}
