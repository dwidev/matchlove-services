package model

import (
	"time"

	"gorm.io/gorm"
)

type TimeField struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *TimeField) BeforeSave(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now()
	return
}

func (t *TimeField) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return
}
