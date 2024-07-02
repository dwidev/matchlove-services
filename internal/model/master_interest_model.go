package model

import "github.com/google/uuid"

type MasterInterestModel struct {
	ID   uuid.UUID `gorm:"primaryKey;type:varchar(36);"`
	Code string    `gorm:"type:varchar(50);unique;not null"`
	Name string    `gorm:"type:varchar(255);"`
}

func (m MasterInterestModel) TableName() string {
	return "master_interest"
}
