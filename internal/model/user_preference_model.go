package model

import (
	"github.com/google/uuid"
)

type UserPreference struct {
	Uuid        uuid.UUID   `gorm:"primaryKey;type:varchar(36);"`
	Account     UserAccount `gorm:"foreignKey:AccountUuid"`
	AccountUuid string      `gorm:"type:varchar(36);uniqueIndex"`

	PreferredGender string `gorm:"type:enum('Male', 'Female', 'Other')"`
	AgeMin          uint8
	AgeMax          uint8
	InterestFor     string `gorm:"size:255"`
	LookingFor      string `gorm:"size:255"`
	Distance        float64
}

func (UserPreference) TableName() string {
	return "user_preference"
}
