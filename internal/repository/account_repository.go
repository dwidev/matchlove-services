package repository

import (
	"matchlove-services/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	UpdateAccountRefreshToken(AccountUuid string, newRefreshToken string) (err error)
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

func (repo *AccountRepository) GetAccountByID(AccountUuid string) (acc *model.UserAccount, err error) {
	where := &model.UserAccount{Uuid: uuid.MustParse(AccountUuid)}
	result := repo.db.Where(where).First(&acc)

	if result.Error != nil {
		panic(result.Error)
	}

	return acc, nil
}

func (repo *AccountRepository) UpdateAccountRefreshToken(AccountUuid string, newRefreshToken string) (err error) {
	account := new(model.UserAccount)
	where := &model.UserAccount{Uuid: uuid.MustParse(AccountUuid)}
	result := repo.db.Model(account).Where(where).Update("refresh_token", newRefreshToken)

	if result.Error != nil {
		panic(result.Error)
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
