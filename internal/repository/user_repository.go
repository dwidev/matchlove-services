package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/helper"
	"matchlove-services/pkg/response"
	"sync"
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
	UpdateProfile(account *model.UserAccount) (*model.UserAccount, error)
	UpdateInterest(accountID string, userInterest []*model.UserInterest) error
	CreateOrUpdateMyLifeStyle(lifestyle *model.UserLifeStyle) error
	CreateOrUpdateMyRoutine(routine *model.UserRoutine) error
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, response.RecordNotFound
	}

	if err != nil {
		logrus.Errorf("GetProfile : error on get user profile %s %s", accountID, err)
		return nil, err
	}

	err = repo.parseNameUserInterest(tx, user.UserInterest)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		logrus.Errorf("GetProfile : error on commit %s", err)
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) UpdateProfile(account *model.UserAccount) (acc *model.UserAccount, err error) {
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Repo UpdateProfile : panic recover txn  %s", r)
			tx.Rollback()
		}
	}()

	// update user profile
	if err = tx.Model(account.UserProfile).
		Select("*").
		Omit("uuid", "account_uuid", "first_name", "last_name", "longitude", "latitude", "age").
		Save(account.UserProfile).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 3)
	wg.Add(3)

	// update user interest
	go func(accountID string, interests []*model.UserInterest) {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("UserRepository.UpdateProfile  : panic recover go update interest  %s", r)
				errChan <- errors.New(fmt.Sprintf("%v", r))
			}
			wg.Done()
		}()

		err := repo.UpdateInterest(accountID, interests)
		if err != nil {
			errChan <- err
		}
	}(account.Uuid.String(), account.UserInterest)

	// update/create user lifestyle
	go func(accountID string, lifestyle *model.UserLifeStyle) {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("UserRepository.UpdateProfile  : panic recover go %s", r)
				errChan <- errors.New(fmt.Sprintf("%v", r))
			}
			wg.Done()
		}()

		if err := repo.CreateOrUpdateMyLifeStyle(lifestyle); err != nil {
			errChan <- err
		}
	}(account.Uuid.String(), account.UserLifeStyle)

	// update/create user routine
	go func(accountID string, routine *model.UserRoutine) {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("UserRepository.UpdateProfile  : panic recover go %s", r)
				errChan <- errors.New(fmt.Sprintf("%v", r))
			}
			wg.Done()
		}()

		if err := repo.CreateOrUpdateMyRoutine(routine); err != nil {
			errChan <- err
		}
	}(account.Uuid.String(), account.UserRoutine)

	wg.Wait()
	select {
	case err := <-errChan:
		return nil, err
	default:
		close(errChan)
		break
	}

	err = tx.Preload("UserPreference").
		Preload("UserProfile").
		Preload("UserInterest").
		Preload("UserLifeStyle").
		Preload("UserRoutine").
		Where("uuid = ?", account.Uuid).
		First(&acc).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, response.RecordNotFound
	}

	err = repo.parseNameUserInterest(tx, acc.UserInterest)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return acc, nil
}

func (repo *UserRepository) parseNameUserInterest(tx *gorm.DB, userInterest []*model.UserInterest) (err error) {
	for i, ui := range userInterest {
		interest := new(model.MasterInterestModel)
		err := tx.Model(interest).Where("code = ?", ui.InterestCode).First(interest).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			if err != nil {
				tx.Rollback()
				logrus.Errorf("error on parse name interest %s", err)
				return err
			}

			userInterest[i].Name = interest.Name
		}
	}

	return nil
}

func (repo *UserRepository) UpdateInterest(accountID string, userInterest []*model.UserInterest) error {
	//panic("ERROR UPDATE INTEREST")
	var existingInterest []*model.UserInterest
	if err := repo.db.Find(&existingInterest, "account_id = ?", accountID).Error; err != nil {
		logrus.Errorf("(UserRepository.UpdateInterest) error on get existingInterest %s", err)
		return err
	}

	existingCond := make(map[string]bool)
	for _, existing := range existingInterest {
		existingCond[existing.InterestCode] = true
	}

	for _, interest := range userInterest {
		if _, exist := existingCond[interest.InterestCode]; exist {
			delete(existingCond, interest.InterestCode)
		} else {
			if err := repo.db.Create(&interest).Error; err != nil {
				logrus.Errorf("(UserRepository.UpdateInterest) error on create interest %s", err)
				return err
			}
		}
	}

	for key := range existingCond {
		if err := repo.db.Where("interest_code = ?", key).Delete(&model.UserInterest{}).Error; err != nil {
			logrus.Errorf("(UserRepository.UpdateInterest) error on delete existing interest %s", err)
			return err
		}
	}

	return nil
}

func (repo *UserRepository) CreateOrUpdateMyLifeStyle(lifestyle *model.UserLifeStyle) error {
	lifestyle.ID = uuid.New()
	if err := repo.db.Save(lifestyle).Error; err != nil {
		logrus.Errorf("(UserRepository.CreateOrUpdateLifeStyle) error on create or update lifestyle %s", err)
		return err
	}

	return nil
}

func (repo *UserRepository) CreateOrUpdateMyRoutine(routine *model.UserRoutine) error {
	routine.ID = uuid.New()
	if err := repo.db.Save(routine).Error; err != nil {
		logrus.Errorf("(UserRepository.CreateOrUpdateMyRoutine) error on create or update routine %s", err)
		return err
	}

	return nil
}
