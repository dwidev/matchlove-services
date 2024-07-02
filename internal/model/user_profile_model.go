package model

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	Uuid        uuid.UUID   `gorm:"primaryKey;type:varchar(36);"`
	Account     UserAccount `gorm:"unique;foreignKey:AccountUuid"`
	AccountUuid string      `gorm:"type:varchar(36);"`

	FirstName         string `gorm:"type:varchar(50)"`
	LastName          string `gorm:"type:varchar(50);"`
	Gender            string `gorm:"type:enum('Male', 'Female', 'Other')"`
	DateOfBirth       time.Time
	Bio               string `gorm:"type:text"`
	ProfilePictureURL string `gorm:"size:255"`
	Longitude         float64
	Latitude          float64
}

func (UserProfile) TableName() string {
	return "user_profile"
}
