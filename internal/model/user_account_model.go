package model

import (
	"time"

	"github.com/google/uuid"
)

type UserAccount struct {
	Uuid              uuid.UUID  `gorm:"primaryKey;type:varchar(36);" json:"uuid"`
	Username          string     `gorm:"unique;type:varchar(255)" json:"username"`
	Email             string     `gorm:"unique;type:varchar(255)" json:"email,omitempty"`
	CreatedAt         *time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	Password          string     `gorm:"type:varchar(255)" json:"-"`
	IsCompleteProfile uint8      `json:"-"`

	LoginActivity []*LoginActivity `gorm:"foreignKey:AccountID" json:"-"`

	UserPreference *UserPreference `gorm:"unique;foreignKey:AccountUuid" json:"user_preference,omitempty"`
	UserProfile    *UserProfile    `gorm:"unique;foreignKey:AccountUuid" json:"user_profile"`
	UserInterest   []*UserInterest `gorm:"foreignKey:AccountID" json:"user_interest"`
	UserLifeStyle  *UserLifeStyle  `gorm:"foreignKey:AccountID" json:"user_life_style"`
	UserRoutine    *UserRoutine    `gorm:"foreignKey:AccountID" json:"user_routine"`
}

func (UserAccount) TableName() string {
	return "user_account"
}
