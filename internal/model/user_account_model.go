package model

import (
	"matchlove-services/internal/dto"
	"time"

	"github.com/google/uuid"
)

type UserAccount struct {
	Uuid              uuid.UUID `gorm:"primaryKey;type:varchar(36);"`
	Username          string    `gorm:"unique;type:varchar(255)"`
	Email             string    `gorm:"unique;type:varchar(255)"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	Password          string    `gorm:"type:varchar(255)"`
	RefreshToken      string
	LastLogin         *time.Time
	IsCompleteProfile uint8
}

func (UserAccount) TableName() string {
	return "user_account"
}

func (a UserAccount) ParseToDTO() *dto.AccountDTO {
	account := &dto.AccountDTO{
		Uuid:     a.Uuid.String(),
		Username: a.Username,
		Email:    a.Email,
	}

	return account
}
