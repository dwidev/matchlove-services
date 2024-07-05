package model

import "github.com/google/uuid"

type UserLifeStyle struct {
	ID                  uuid.UUID `gorm:"primaryKey;type:varchar(36);"`
	AccountID           string    `gorm:"type:varchar(36);" json:"account_uuid"`
	SocialMediaActivity string    `gorm:"type:varchar(255);" json:"social_media_activity"`
	Pets                string    `gorm:"type:varchar(255);" json:"pets"`
	Drinking            string    `gorm:"type:varchar(255);" json:"drinking"`
	Smoking             string    `gorm:"type:varchar(255);" json:"smoking"`
	Workout             string    `gorm:"type:varchar(255);" json:"workout"`
	SleepingHabits      string    `gorm:"type:varchar(255);" json:"sleepingHabits"`
}

func (l UserLifeStyle) TableName() string {
	return "user_lifestyle"
}
