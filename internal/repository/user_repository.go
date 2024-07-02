package repository

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/response"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewUserRepository(db *gorm.DB, ar IAccountRepository) IUserRepository {
	return &UserRepository{
		db:                db,
		AccountRepository: ar,
	}
}

type IUserRepository interface {
	RegisterUser(dto *dto.UserProfileRegisterDTO) error
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
			log.Println("recover on RegisterUser", r)
			tx.Rollback()
		}
	}()

	userProfile := model.UserProfile{
		Uuid:        uuid.New(),
		AccountUuid: dto.AccountId,
		FirstName:   dto.Name,
		LastName:    "",
		Gender:      dto.Gender,
		DateOfBirth: *parseTime(dto.DOB),
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

func parseTime(dateTime string) *time.Time {
	layout := "2006-01-02 15:04:05"

	// Parse the date-time string into a time.Time object in the specified location
	parsedTime, err := time.Parse(layout, dateTime)
	if err != nil {
		fmt.Println("Error parsing date-time with location:", err)
		panic(err)
	}

	return &parsedTime
}
