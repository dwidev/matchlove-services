package model

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	Uuid        uuid.UUID `gorm:"primaryKey;type:varchar(36);" json:"uuid"`
	AccountUuid string    `gorm:"type:varchar(36);" json:"account_uuid"`

	ProfilePictureURL string    `gorm:"size:255" json:"profile_picture_url"`
	FirstName         string    `gorm:"type:varchar(50)" json:"first_name"`
	LastName          string    `gorm:"type:varchar(50);" json:"last_name"`
	Bio               string    `gorm:"type:text" json:"bio"`
	Gender            string    `gorm:"type:enum('Male', 'Female', 'Other')" json:"gender"`
	DateOfBirth       time.Time `gorm:"type:timestamp;" json:"date_of_birth"`
	Age               int       `gorm:"type:int;" json:"age"`
	Height            float64   `gorm:"type:float;" json:"height"`
	Weight            float64   `gorm:"type:float;" json:"weight"`
	Zodiac            string    `gorm:"type:text" json:"zodiac"`
	BloodType         string    `gorm:"type:varchar(50)" json:"blood_type"`
	Education         string    `gorm:"type:varchar(50)" json:"education"`
	Personality       string    `gorm:"type:varchar(50)" json:"personality"`
	LookingFor        string    `gorm:"type:varchar(50)" json:"looking_for"`
	LoveLanguage      string    `gorm:"type:varchar(50)" json:"love_language"`
	City              string    `gorm:"type:varchar(50)" json:"city"`
	Longitude         float64   `json:"longitude"`
	Latitude          float64   `json:"latitude"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}
