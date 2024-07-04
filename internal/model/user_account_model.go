package model

import (
	"time"

	"github.com/google/uuid"
)

type UserAccount struct {
	Uuid              uuid.UUID  `gorm:"primaryKey;type:varchar(36);" json:"uuid"`
	Username          string     `gorm:"unique;type:varchar(255)" json:"username"`
	Email             string     `gorm:"unique;type:varchar(255)" json:"email"`
	CreatedAt         time.Time  `gorm:"autoCreateTime" json:"created_at"`
	Password          string     `gorm:"type:varchar(255)" json:"-"`
	RefreshToken      string     `json:"-"`
	LastLogin         *time.Time `json:"last_login,omitempty"`
	IsCompleteProfile uint8      `json:"-"`

	UserPreference UserPreference `gorm:"unique;foreignKey:AccountUuid" json:"user_preference"`
	UserProfile    UserProfile    `gorm:"unique;foreignKey:AccountUuid" json:"user_profile"`
	UserInterest   []UserInterest `gorm:"foreignKey:AccountID" json:"user_interest"`
}

func (UserAccount) TableName() string {
	return "user_account"
}
