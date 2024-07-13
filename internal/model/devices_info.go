package model

import (
	"github.com/google/uuid"
)

type DevicesInfo struct {
	ID              uuid.UUID `gorm:"primary_key;"`
	LoginActivityID string    `gorm:"type:varchar(36);"`
	Imei            string
	Platform        string
	OsVersion       string
	DeviceName      string
}

func (DevicesInfo) TableName() string {
	return "devices_info"
}
