package model

type Likes struct {
	ID      uint64 `gorm:"primary_key;auto_increment" json:"-"`
	LikerID string `gorm:"type:varchar(36)"`
	LikedID string `gorm:"type:varchar(36);uniqueIndex"`

	FirstUser  *UserAccount `gorm:"foreignKey:LikerID"`
	SecondUser *UserAccount `gorm:"foreignKey:LikedID"`

	TimeField
}

func (m Likes) TableName() string {
	return "likes"
}
