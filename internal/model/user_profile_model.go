package model

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	Uuid        uuid.UUID `gorm:"primaryKey;type:varchar(36);" json:"uuid"`
	AccountUuid string    `gorm:"type:varchar(36);" json:"account_uuid"`

	FirstName         string    `gorm:"type:varchar(50)" json:"first_name"`
	LastName          string    `gorm:"type:varchar(50);" json:"last_name"`
	Gender            string    `gorm:"type:enum('Male', 'Female', 'Other')" json:"gender"`
	DateOfBirth       time.Time `gorm:"type:timestamp;" json:"date_of_birth"`
	Age               int       `gorm:"type:int;" json:"age"`
	Bio               string    `gorm:"type:text" json:"bio"`
	ProfilePictureURL string    `gorm:"size:255" json:"profile_picture_url"`
	Longitude         float64   `json:"longitude"`
	Latitude          float64   `json:"latitude"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}
