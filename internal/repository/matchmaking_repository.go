package repository

import (
	"fmt"
	"gorm.io/gorm"
	"matchlove-services/internal/dto"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/cache"
	"strings"
	"time"
)

func NewMatchmakingRepository(db *gorm.DB, cache cache.Cache) IMatchmakingRepository {
	return &MatchmakingRepository{
		db:    db,
		cache: cache,
	}
}

type IMatchmakingRepository interface {
	GetMatchSuggestions(dto *dto.MatchSuggestionsRequestDto) (*dto.PaginationResultDTO, error)
	CreateLike(dto *dto.LikeRequestDTO) (liked bool, err error)
	CheckForMatches(dto *dto.LikeRequestDTO) (matches bool, err error)
	CreateMatches(dto *dto.LikeRequestDTO) (bool, error)
}

type MatchmakingRepository struct {
	db    *gorm.DB
	cache cache.Cache
}

func (repo *MatchmakingRepository) CalculateTotalUser(accountID string, preferenceGender string) (*int64, error) {
	var count int64
	query := repo.db.Model(&model.UserProfile{}).Where("account_uuid != ?", accountID)
	query.Where("gender = ?", preferenceGender).Count(&count)

	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
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

	if !cfg.disableDistance && preference.Distance != 0 {
		query.Where(fmt.Sprintf("%s <= ?", distance), preference.Distance)
	}

	if preference.PreferredGender != "" {
		query.Where("UserProfile.gender = ?", preference.PreferredGender)
	}

	if !cfg.disableAge {
		if preference.AgeMin != 0 || preference.AgeMax != 0 {
			query.Where("UserProfile.age >= ?", preference.AgeMin).
				Where("UserProfile.age <= ?", preference.AgeMax)
		} else {
			query.Where("UserProfile.age >= ?", profile.Age).
				Where("UserProfile.age <= ?", profile.Age+10)
		}
	}

	if !cfg.disableLookingFor && preference.LookingFor != "" {
		query.Where("UserProfile.looking_for = ?", preference.LookingFor)
	}

	if !cfg.disableInterest {
		interest := strings.Split(preference.InterestFor, "#")
		if len(interest) != 0 {
			query.Preload("UserInterest", "interest_code IN (?)", interest)
		}
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

func (repo *MatchmakingRepository) GetMatchSuggestions(request *dto.MatchSuggestionsRequestDto) (*dto.PaginationResultDTO, error) {
	perPage := request.PerPage

	profile := new(model.UserProfile)
	if err := repo.db.Where("account_uuid = ?", request.AccountID).First(profile).Error; err != nil {
		// TODO(): create get list match without preference user
		return nil, err
	}

	preference := new(model.UserPreference)
	if err := repo.db.Where("account_uuid = ?", request.AccountID).First(preference).Error; err != nil {
		// TODO(): create get list match without preference user
		return nil, err
	}

	total, err := repo.CalculateTotalUser(request.AccountID, preference.PreferredGender)
	if err != nil {
		return nil, err
	}

	result := make([]*model.UserAccount, 0)
	engineConfig := &matchMakingEngineConfig{
		profile:    profile,
		preference: preference,
		dto:        request,
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

	if len(result) < perPage {
		engineConfig.dto.PerPage = perPage - len(result)
		engineConfig.NotIncludeInterest()
		notIncludeLookingFor, err := repo.matchMakingEngine(engineConfig)
		if err != nil {
			return nil, err
		}

		result = append(result, notIncludeLookingFor...)
	}

	if len(result) < perPage {
		engineConfig.dto.PerPage = perPage - len(result)
		engineConfig.DisablePreference()
		notIncludeLookingFor, err := repo.matchMakingEngine(engineConfig)
		if err != nil {
			return nil, err
		}

		result = append(result, notIncludeLookingFor...)
	}

	totalPage := (int(*total) + perPage - 1) / perPage
	return &dto.PaginationResultDTO{
		CurrentPage: request.Page,
		TotalData:   int(*total),
		TotalPage:   totalPage,
		Data:        result,
	}, nil
}

func (repo *MatchmakingRepository) CreateLike(dto *dto.LikeRequestDTO) (liked bool, err error) {
	var count int64
	err = repo.db.Model(model.Likes{}).Where("liker_id = ?", dto.FirstUserAccountID).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	likerID := dto.FirstUserAccountID
	likedID := dto.SecondUserAccountID

	likesModel := model.Likes{
		LikerID: likerID,
		LikedID: likedID,
		TimeField: model.TimeField{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := repo.db.Create(&likesModel).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (repo *MatchmakingRepository) CheckForMatches(dto *dto.LikeRequestDTO) (matches bool, err error) {
	var count int64
	err = repo.db.Model(model.Likes{}).Where("liked_id = ?", dto.FirstUserAccountID).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, err
}

func (repo *MatchmakingRepository) CreateMatches(dto *dto.LikeRequestDTO) (bool, error) {
	var count int64
	err := repo.db.Model(model.Matches{}).
		Where("first_user_id = ? OR second_user_id = ?", dto.FirstUserAccountID, dto.FirstUserAccountID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	matchModel := model.Matches{
		FirstUserID:       dto.FirstUserAccountID,
		SecondUserID:      dto.SecondUserAccountID,
		Score:             0,
		MatchAtMobileTime: time.Now(),
		TimeField: model.TimeField{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if count > 0 {
		if err := repo.db.Where("first_user_id = ?", dto.FirstUserAccountID).Updates(&matchModel).Error; err != nil {
			return false, err
		}
		return true, nil
	} else {
		if err := repo.db.Create(&matchModel).Error; err != nil {
			return false, err
		}
		return true, nil
	}
}
