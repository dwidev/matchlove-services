package model

import (
	"time"
)

type Matches struct {
	ID           uint64 `gorm:"primary_key;auto_increment" json:"-"`
	FirstUserID  string `gorm:"type:varchar(36)" json:"first_user_profile_id"`
	SecondUserID string `gorm:"type:varchar(36)" json:"second_user_profile_id"`
	Score        int

	FirstUser  *UserAccount `gorm:"foreignKey:FirstUserID" json:"-"`
	SecondUser *UserAccount `gorm:"foreignKey:SecondUserID" json:"-"`

	MatchAtMobileTime time.Time `json:"match_at_mobile_time"`
	TimeField
}

func (m Matches) TableName() string {
	return "matches"
}
