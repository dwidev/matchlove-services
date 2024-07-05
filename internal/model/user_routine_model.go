package model

import "github.com/google/uuid"

type UserRoutine struct {
	ID             uuid.UUID `gorm:"primaryKey;type:varchar(36);"`
	AccountID      string    `gorm:"type:varchar(36);" json:"account_uuid"`
	MyWeekend      string    `gorm:"type:varchar(255);"`
	MyHangouts     string    `gorm:"type:varchar(255);"`
	MorningRoutine string    `gorm:"type:varchar(255);"`
}

func (u UserRoutine) TableName() string {
	return "user_routine"
}
