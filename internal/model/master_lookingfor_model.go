package model

type MasterLookingFor struct {
	ID   uint64 `gorm:"primaryKey;auto_increment"`
	Code string `gorm:"type:varchar(50);unique;not null"`
	Name string `gorm:"type:varchar(255);"`
}

func (m MasterLookingFor) TableName() string {
	return "master_looking_for"
}
