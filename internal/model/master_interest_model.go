package model

type MasterInterestModel struct {
	ID   uint64 `gorm:"primary_key;auto_increment"`
	Code string `gorm:"type:varchar(50);unique;not null"`
	Name string `gorm:"type:varchar(255);"`
}

func (m MasterInterestModel) TableName() string {
	return "master_interest"
}
