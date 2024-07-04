package model

type UserInterest struct {
	ID           int64  `gorm:"primary_key;AUTO_INCREMENT:true" json:"-"`
	AccountID    string `gorm:"type:varchar(36);" json:"-"`
	InterestCode string `gorm:"varchar(50)" json:"code"`
	Name         string `json:"name"`
}

func (i UserInterest) TableName() string {
	return "user_interest"
}
