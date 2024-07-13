package model

import (
	"github.com/google/uuid"
	"time"
)

func NewLoginActivity() *LoginActivity {
	return &LoginActivity{
		ID: uuid.New(),
	}
}

type LoginActivity struct {
	ID           uuid.UUID  `gorm:"primary_key;type:varchar(36);"`
	AccountID    string     `gorm:"type:varchar(36);" json:"-"`
	AccessToken  string     `json:"-"`
	RefreshToken string     `json:"-"`
	LoginAt      *time.Time `json:"-"`
	LastLogin    *time.Time `json:"last_login,omitempty"`

	DevicesInfo DevicesInfo `gorm:"unique;foreignKey:LoginActivityID"`
}

func (l *LoginActivity) TableName() string {
	return "login_activities"
}
func (l *LoginActivity) SetLoginAt() {
	now := time.Now()
	l.LoginAt = &now
}

func (l *LoginActivity) BeforeSave() error {
	l.SetLoginAt()
	return nil
}

func (l *LoginActivity) BeforeUpdate() error {
	l.SetLoginAt()
	return nil
}
