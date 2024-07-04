package repository

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/helper"
	"matchlove-services/pkg/response"
)

func NewUserRepository(db *gorm.DB, ar IAccountRepository) IUserRepository {
	return &UserRepository{
		db:                db,
		AccountRepository: ar,
	}
}

type IUserRepository interface {
	RegisterUser(dto *dto.UserProfileRegisterDTO) error
	GetProfile(accountID string) (*model.UserAccount, error)
}

type UserRepository struct {
	db                *gorm.DB
	AccountRepository IAccountRepository
}

func (repo *UserRepository) RegisterUser(dto *dto.UserProfileRegisterDTO) error {
	// start transaction
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic recover on RegisterUser", r)
			tx.Rollback()
		}
	}()

	userProfile := model.UserProfile{
		Uuid:        uuid.New(),
		AccountUuid: dto.AccountId,
		FirstName:   dto.Name,
		LastName:    "",
		Gender:      dto.Gender,
		DateOfBirth: helper.ParseDateTime(dto.DOB),
		Bio:         "",
		Longitude:   dto.Longitude,
		Latitude:    dto.Latitude,
	}

	// check user profile already exist
	var exist int64
	if err := tx.Model(&model.UserProfile{}).Where(&model.UserProfile{AccountUuid: dto.AccountId}).Count(&exist).Error; err != nil {
		logrus.Errorf("error on check user profile exist %s", err)
		tx.Rollback()
		return err
	}

	if exist > 0 {
		tx.Rollback()
		return response.AlreadyExist
	}

	// create user profile
	if err := tx.Create(&userProfile).Error; err != nil {
		logrus.Errorf("error on create user profile %s", err)
		tx.Rollback()
		return err
	}

	// insert user interest
	for _, code := range dto.InterestFor {
		i := &model.UserInterest{
			AccountID:    dto.AccountId,
			InterestCode: code,
		}
		if err := tx.Create(i).Error; err != nil {
			logrus.Errorf("error on create user interest %s", err)
			tx.Rollback()
			return err
		}
	}

	// create/update user preference
	userPreference := model.UserPreference{
		Uuid:            uuid.New(),
		AccountUuid:     dto.AccountId,
		AgeMin:          dto.Age - 3,
		AgeMax:          dto.Age + 3,
		PreferredGender: dto.ToPreferedGender(),
		Distance:        float64(dto.Distance),
		LookingFor:      dto.LookingFor,
		InterestFor:     dto.JoinInterest(),
	}

	tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "account_uuid"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"preferred_gender",
			"age_min",
			"age_max",
			"interest_for",
			"looking_for",
			"distance",
			"latitude",
			"longitude",
		}),
	}).Create(&userPreference)

	// update complete profile at user account
	account := new(model.UserAccount)
	where := &model.UserAccount{Uuid: uuid.MustParse(dto.AccountId)}
	if err := tx.Model(account).Where(where).Update("is_complete_profile", 1).Error; err != nil {
		logrus.Errorf("Error update complete profile %s", err)
		tx.Rollback()
		return err
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		logrus.Errorf("Error on transaction Register User %s", err)
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *UserRepository) GetProfile(accountID string) (*model.UserAccount, error) {
	user := new(model.UserAccount)
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("GetProfile : panic recover txn  %s", r)
			tx.Rollback()
		}
	}()

	err := tx.Preload("UserPreference").
		Preload("UserProfile").
		Preload("UserInterest").
		Where("uuid = ?", accountID).
		First(user).Error
	if err != nil {
		logrus.Errorf("GetProfile : error on get user profile %s %s", accountID, err)
		return nil, err
	}

	for i, ui := range user.UserInterest {
		interest := new(model.MasterInterestModel)
		err := tx.Model(interest).Where("code = ?", ui.InterestCode).First(interest).Error
		if err != nil {
			tx.Rollback()
			logrus.Errorf("GetProfile : error on get interest %s", err)
			return nil, err
		}

		user.UserInterest[i].Name = interest.Name
	}

	err = tx.Commit().Error
	if err != nil {
		logrus.Errorf("GetProfile : error on commit %s", err)
		return nil, err
	}

	return user, nil
}
