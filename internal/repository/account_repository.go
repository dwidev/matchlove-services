package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"matchlove-services/internal/model"
)

type IAccountRepository interface {
	CreateNewAccount(account *model.UserAccount) (acc *model.UserAccount, err error)
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

func (repo *AccountRepository) ChangePassword(AccountUuid string, newPassword string) error {
	account := new(model.UserAccount)
	results := repo.db.Model(account).Where("uuid = ?", AccountUuid).Update("password", newPassword)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
