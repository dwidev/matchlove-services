package seeder

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/helper"
	"strings"
	"time"
)

func SeedUsers(db *gorm.DB) error {
	userAccountData = append(userAccountData, userAccountData...)
	userProfileData = append(userProfileData, userProfileComplete...)
	userPreferenceData = append(userPreferenceData, userPreferenceData...)

	db.Exec("DELETE FROM user_interest")
	db.Exec("DELETE FROM user_preference")
	db.Exec("DELETE FROM user_profile")
	db.Exec("DELETE FROM user_account")

	tx := db.Begin()
	var interest []model.MasterInterestModel
	if err := tx.Find(&interest).Error; err != nil {
		tx.Rollback()
	}

	for i, account := range userAccountData {
		pass, err := helper.ToHash(fmt.Sprintf("userdummy%d", i))
		if err != nil {
			tx.Rollback()
			return err
		}

		account.Uuid = uuid.New()
		account.Username = fmt.Sprintf("%s%d", account.Username, i)
		account.Email = fmt.Sprintf("userdummy%d@gmail.com", i)
		account.Password = pass

		profile := userProfileData[i]
		if profile.Height == 0 {
			profile.Height = 189
		}
		if profile.Weight == 0 {
			profile.Weight = 74
		}
		if profile.LookingFor == "" {
			profile.LookingFor = "LT_PARTNER"
		}
		if profile.Zodiac == "" {
			profile.Zodiac = "Pisces"
		}
		if profile.BloodType == "" {
			profile.BloodType = "AB"
		}
		if profile.Education == "" {
			profile.Education = "High School"
		}
		if profile.Personality == "" {
			profile.Personality = "EXTROVERT"
		}
		if profile.LoveLanguage == "" {
			profile.LoveLanguage = "QUALITY_TIME"
		}
		if profile.City == "" {
			profile.City = "BOGOR"
		}

		profile.Uuid = uuid.New()
		profile.AccountUuid = account.Uuid.String()
		r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
		profile.Age = r.Intn(20) + 17

		preference := userPreferenceData[i]
		preference.Uuid = uuid.New()
		preference.AccountUuid = account.Uuid.String()
		if profile.Gender == "Male" {
			preference.PreferredGender = "Female"
		} else {
			preference.PreferredGender = "Male"
		}

		interestRandom := helper.RandomArray(interest, 2)
		var interestCode []string
		for _, i := range interestRandom {
			interestCode = append(interestCode, i.Code)
		}
		preference.InterestFor = strings.Join(interestCode, "#")
		preference.LookingFor = profile.LookingFor

		if err := tx.Create(&account).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Create(&profile).Error; err != nil {
			tx.Rollback()
			return err
		}

		for _, code := range interestCode {
			interest := model.UserInterest{
				AccountID:    account.Uuid.String(),
				InterestCode: code,
			}
			if interest.InterestCode == "" {
				interest.InterestCode = "GAMING"
			}

			if err := tx.Create(&interest).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		if err := tx.Create(&preference).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
