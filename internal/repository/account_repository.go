package repository

import (
	"errors"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/jwt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	CreateNewAccount(account *model.UserAccount) (acc *model.UserAccount, err error)
	RecordLoginActivity(recordLogin *dto.RecordLoginActivityDto, token jwt.TokenPayload) (err error)
	GetAccountByID(AccountUuid string) (acc *model.UserAccount, err error)
	ChangePassword(AccountUuid string, newPassword string) error
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &AccountRepository{
		db: db,
	}
}

type AccountRepository struct {
	db *gorm.DB
}

func (repo *AccountRepository) CreateNewAccount(account *model.UserAccount) (*model.UserAccount, error) {
	if err := repo.db.Create(&account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

func (repo *AccountRepository) GetAccountByID(AccountUuid string) (acc *model.UserAccount, err error) {
	where := &model.UserAccount{Uuid: uuid.MustParse(AccountUuid)}
	result := repo.db.Where(where).First(&acc)

	if result.Error != nil {
		panic(result.Error)
	}

	return acc, nil
}

func (repo *AccountRepository) RecordLoginActivity(recordLogin *dto.RecordLoginActivityDto, token jwt.TokenPayload) (err error) {
	imei := recordLogin.DevicesInfo.Imei
	platform := recordLogin.DevicesInfo.Platform

	loginActivity := recordLogin.LoginActivity

	var exist int64
	if err = repo.db.Model(&model.LoginActivity{}).Joins("DevicesInfo").
		Where("account_id = ?", loginActivity.AccountID).
		Where("imei = ? AND platform = ?", imei, platform).
		Count(&exist).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exist = 0
		} else {
			return err
		}
	}

	if exist != 0 {
		updates := model.LoginActivity{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		}
		sub := repo.db.Model(&model.DevicesInfo{}).Select("login_activity_id").Where("imei = ? AND platform = ?", imei, platform)
		if err = repo.db.Model(&model.LoginActivity{}).Where("ID = (?)", sub).Updates(updates).Error; err != nil {
			return err
		}
	} else {
		now := time.Now()
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

func (repo *AccountRepository) ChangePassword(AccountUuid string, newPassword string) error {
	account := new(model.UserAccount)
	results := repo.db.Model(account).Where("uuid = ?", AccountUuid).Update("password", newPassword)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
