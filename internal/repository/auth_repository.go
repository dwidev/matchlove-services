package repository

import (
	"errors"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/helper"
	"matchlove-services/pkg/response"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	LoginWithEmail(email string) (acc *model.UserAccount, err error)
	LoginWithPassword(req *dto.LoginPassRequestDTO) (acc *model.UserAccount, err error)
	Logout(accountId string) error
	ValidateRefreshToken(refreshToken string) (res bool, err error)
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}

type AuthRepository struct {
	db *gorm.DB
}

func (repo *AuthRepository) LoginWithEmail(email string) (*model.UserAccount, error) {
	account := new(model.UserAccount)
	where := &model.UserAccount{Email: email}
	err := repo.db.Where(where).First(account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return account, err
}

func (repo *AuthRepository) LoginWithPassword(req *dto.LoginPassRequestDTO) (*model.UserAccount, error) {
	account, err := repo.LoginWithEmail(req.Email)
	if err != nil {
		return nil, err
	}

	valid, err := helper.ToCompare(req.Password, account.Password)
	if !valid && err != nil {
		return nil, response.PassNoValid
	}

	return account, nil
}

func (repo *AuthRepository) Logout(accountId string) error {
	err := repo.UpdateAccountRefreshToken(accountId, "")
	if err != nil {
		panic(err)
	}

	return nil
}

func (repo *AuthRepository) GetAccountByUserID(accountId string) (*model.UserAccount, error) {
	account := new(model.UserAccount)
	where := &model.UserAccount{Uuid: uuid.MustParse(accountId)}
	result := repo.db.Where(where).First(&account)

	if result.Error != nil {
		panic(result.Error)
	}

	return account, nil
}

func (repo *AuthRepository) UpdateAccountRefreshToken(accountId string, newRefreshToken string) error {
	account := new(model.UserAccount)
	where := &model.UserAccount{Uuid: uuid.MustParse(accountId)}
	result := repo.db.Model(account).Where(where).Update("refresh_token", newRefreshToken)

	if result.Error != nil {
		panic(result.Error)
	}

	return nil
}

func (repo *AuthRepository) ValidateRefreshToken(refreshToken string) (res bool, err error) {
	var count int64
	err = repo.db.Model(&model.LoginActivity{}).Where("refresh_token = ?", refreshToken).Count(&count).Error
	if err != nil {
		return false, err
	}

	res = count > 0
	return res, nil
}
