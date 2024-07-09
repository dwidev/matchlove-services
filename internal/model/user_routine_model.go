package model

import "github.com/google/uuid"

type UserRoutine struct {
	ID             uuid.UUID `gorm:"primaryKey;type:varchar(36);" json:"id"`
	AccountID      string    `gorm:"type:varchar(36);uniqueIndex" json:"account_uuid"`
	MyWeekend      string    `gorm:"type:varchar(255);" json:"my_weekend"`
	MyHangouts     string    `gorm:"type:varchar(255);" json:"my_hangouts"`
	MorningRoutine string    `gorm:"type:varchar(255);" json:"morning_routine"`
}

func (u UserRoutine) TableName() string {
	return "user_routine"
}
