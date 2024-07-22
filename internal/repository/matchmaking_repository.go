package repository

import (
	"fmt"
	"gorm.io/gorm"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"strings"
)

func NewMatchmakingRepository(db *gorm.DB) IMatchmakingRepository {
	return &MatchmakingRepository{
		db: db,
	}
}

type IMatchmakingRepository interface {
	GetMatchSuggestions(dto *dto.MatchSuggestionsRequestDto) ([]*model.UserAccount, error)
}

type MatchmakingRepository struct {
	db *gorm.DB
}

func (repo *MatchmakingRepository) matchMakingEngine(cfg *matchMakingEngineConfig) ([]*model.UserAccount, error) {
	cfgEngine := makeConfig(cfg)

	profile := cfgEngine.profile
	preference := cfgEngine.preference
	accountID := profile.AccountUuid

	offset := cfgEngine.Offset()
	limit := cfgEngine.Limit()

	distance := fmt.Sprintf(`ST_Distance_Sphere(point(UserProfile.longitude, UserProfile.latitude),point(%f, %f)) / 1000`, profile.Longitude, profile.Latitude)
	query := repo.db.
		Table("user_account").
		Select(fmt.Sprintf("%s AS distance, user_account.*", distance)).
		Omit("email").
		Joins("UserProfile").
		Joins("UserLifeStyle").
		Joins("UserRoutine").
		Where("user_account.uuid != ?", accountID)
	if preference.Distance != 0 {
		query.Where(fmt.Sprintf("%s <= ?", distance), preference.Distance)
	}

	if preference.PreferredGender != "" {
		query.Where("UserProfile.gender = ?", preference.PreferredGender)
	}

	if preference.AgeMin != 0 || preference.AgeMax != 0 {
		query.Where("UserProfile.age >= ?", preference.AgeMin).
			Where("UserProfile.age <= ?", preference.AgeMax)
	} else {
		query.Where("UserProfile.age >= ?", profile.Age).
			Where("UserProfile.age <= ?", profile.Age+10)
	}

	if preference.LookingFor != "" {
		query.Where("UserProfile.looking_for = ?", preference.LookingFor)
	}

	interest := strings.Split(preference.InterestFor, "#")
	if len(interest) != 0 {
		query.Preload("UserInterest", "interest_code IN (?)", interest)
	}

	sameData, ok := cfgEngine.SameData()
	if ok {
		query = query.Where("user_account.uuid NOT IN (?)", sameData)
	}

	resultFromQuery := make([]*model.UserAccount, 0)
	if err := query.Offset(offset).Limit(limit).
		Find(&resultFromQuery).
		Order("user_account.uuid").Error; err != nil {
		return nil, err
	}

	result := make([]*model.UserAccount, 0)

	for _, user := range resultFromQuery {
		cfgEngine.existUser[user.Uuid.String()] = true
		result = append(result, user)
	}

	return result, nil
}

func (repo *MatchmakingRepository) GetMatchSuggestions(dto *dto.MatchSuggestionsRequestDto) ([]*model.UserAccount, error) {
	perPage := dto.PerPage
	profile := new(model.UserProfile)
	if err := repo.db.Where("account_uuid = ?", dto.AccountID).First(profile).Error; err != nil {
		// TODO(): create get list match without preference user
		return nil, err
	}

	preference := new(model.UserPreference)
	if err := repo.db.Where("account_uuid = ?", dto.AccountID).First(preference).Error; err != nil {
		// TODO(): create get list match without preference user
		return nil, err
	}

	result := make([]*model.UserAccount, 0)
	engineConfig := &matchMakingEngineConfig{
		profile:    profile,
		preference: preference,
		dto:        dto,
	}
	dataByDistance, err := repo.matchMakingEngine(engineConfig)
	if err != nil {
		return nil, err
	}
	result = append(result, dataByDistance...)

	if len(result) < perPage {
		engineConfig.dto.PerPage = perPage - len(result)
		engineConfig.ExpandDistance()
		upRangeDistance, err := repo.matchMakingEngine(engineConfig)
		if err != nil {
			return nil, err
		}

		result = append(result, upRangeDistance...)
	}

	if len(result) < perPage {
		engineConfig.dto.PerPage = perPage - len(result)
		engineConfig.ExpandAgeMinAndMax()
		expandAge, err := repo.matchMakingEngine(engineConfig)
		if err != nil {
			return nil, err
		}

		result = append(result, expandAge...)
	}

	if len(result) < perPage {
		engineConfig.dto.PerPage = perPage - len(result)
		engineConfig.NotIncludeLookingFor()
		notIncludeLookingFor, err := repo.matchMakingEngine(engineConfig)
		if err != nil {
			return nil, err
		}

		result = append(result, notIncludeLookingFor...)
	}

	return result, nil
}
