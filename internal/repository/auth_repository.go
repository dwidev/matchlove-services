package repository

import (
	"errors"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/helper"
	"matchlove-services/pkg/jwt"
	"matchlove-services/pkg/response"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	LoginWithEmail(email string) (acc *model.UserAccount, err error)
	LoginWithPassword(req *dto.LoginPassRequestDTO) (acc *model.UserAccount, err error)
	Logout(accountId string, info *model.DevicesInfo) error
	ValidateRefreshToken(refreshToken string) (res bool, err error)
	UpdateRefreshToken(oldToken string, newToken string) (err error)
	RecordLoginActivity(recordLogin *dto.RecordLoginActivityDto, token jwt.TokenPayload) (err error)
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

func (repo *AuthRepository) Logout(accountId string, info *model.DevicesInfo) error {
	loginActivity := new(model.LoginActivity)
	if err := repo.db.Model(loginActivity).
		Joins("DevicesInfo").
		Where("account_id", accountId).
		Where("imei = ? AND platform = ?", info.Imei, info.Platform).
		First(loginActivity).Error; err != nil {
		return err
	}

	updates := map[string]interface{}{
		"refresh_token": "",
		"access_token":  "",
		"last_login":    time.Now(),
	}
	if err := repo.db.Model(loginActivity).Where("id = ?", loginActivity.ID).Updates(updates).Error; err != nil {
		return err
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

func (repo *AuthRepository) ValidateRefreshToken(refreshToken string) (res bool, err error) {
	var count int64
	err = repo.db.Model(&model.LoginActivity{}).Where("refresh_token = ?", refreshToken).Count(&count).Error
	if err != nil {
		return false, err
	}

	res = count > 0
	return res, nil
}

func (repo *AuthRepository) UpdateRefreshToken(oldToken string, newToken string) (err error) {
	err = repo.db.Model(&model.LoginActivity{}).Where("refresh_token = ?", oldToken).Update("refresh_token", newToken).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *AuthRepository) RecordLoginActivity(recordLogin *dto.RecordLoginActivityDto, token jwt.TokenPayload) (err error) {
	imei := recordLogin.DevicesInfo.Imei
	platform := recordLogin.DevicesInfo.Platform
	AccountID := recordLogin.LoginActivity.AccountID

	loginActivity := new(model.LoginActivity)

	var exist int64
	if err = repo.db.Model(loginActivity).Joins("DevicesInfo").
		Where("account_id = ?", AccountID).
		Where("imei = ? AND platform = ?", imei, platform).
		First(loginActivity).
		Count(&exist).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exist = 0
		} else {
			return err
		}
	}

	if exist != 0 {
		now := time.Now()
		updates := model.LoginActivity{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			LoginAt:      &now,
		}

		if err = repo.db.Model(&model.LoginActivity{}).Where("ID = (?)", loginActivity.ID).Updates(updates).Error; err != nil {
			return err
		}
	} else {
		now := time.Now()
		loginActivity := recordLogin.LoginActivity
		loginActivity.LoginAt = &now
		loginActivity.AccessToken = token.AccessToken
		loginActivity.RefreshToken = token.RefreshToken
		if err = repo.db.Create(&loginActivity).Error; err != nil {
			return err
		}

		deviceInfo := recordLogin.DevicesInfo
		deviceInfo.LoginActivityID = loginActivity.ID.String()
		if err = repo.db.Create(deviceInfo).Error; err != nil {
			return err
		}
	}

	return nil
}
