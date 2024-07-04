package model

type UserInterestModel struct {
	ID           int64       `gorm:"primary_key;AUTO_INCREMENT:true" json:"id"`
	Account      UserAccount `gorm:"foreignKey:AccountID" json:"account"`
	AccountID    string      `gorm:"type:varchar(36);" json:"account_id"`
	InterestCode string      `gorm:"varchar(50)" json:"interest_code"`
}

func (i UserInterestModel) TableName() string {
	return "user_interest"
}
