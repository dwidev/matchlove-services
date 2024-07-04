package model

type MasterInterestModel struct {
	ID   uint64 `gorm:"primary_key;auto_increment" json:"-"`
	Code string `gorm:"type:varchar(50);unique;not null" json:"code"`
	Name string `gorm:"type:varchar(255);" json:"name"`
}

func (m MasterInterestModel) TableName() string {
	return "master_interest"
}
