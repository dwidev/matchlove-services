package model

import (
	"github.com/google/uuid"
)

type UserPreference struct {
	Uuid        uuid.UUID `gorm:"primaryKey;type:varchar(36);" json:"uuid"`
	AccountUuid string    `gorm:"type:varchar(36);uniqueIndex" json:"account_uuid"`

	PreferredGender string  `gorm:"type:enum('Male', 'Female', 'Other')" json:"preferred_gender"`
	AgeMin          uint8   `gorm:"type:int" json:"age_min"`
	AgeMax          uint8   `gorm:"type:int" json:"age_max"`
	InterestFor     string  `gorm:"size:255" json:"-"`
	LookingFor      string  `gorm:"size:255" json:"looking_for"`
	Distance        float64 `json:"distance"`
}

func (UserPreference) TableName() string {
	return "user_preference"
}
