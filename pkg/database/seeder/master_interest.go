package seeder

import (
	"gorm.io/gorm"
	"matchlove-services/internal/model"
)

func SeedMasterInterest(db *gorm.DB) error {
	data := []map[string]string{
		{
			"name": "Traveling",
			"code": "TRAVELING",
		},
		{
			"name": "Foodie",
			"code": "FOODIE",
		},
		{
			"name": "Outdoor Adventures",
			"code": "OUTDOOR_ADVENTURES",
		},
		{
			"name": "Fitness",
			"code": "FITNESS",
		},
		{
			"name": "Movies",
			"code": "MOVIES",
		},
		{
			"name": "Music/Spotify",
			"code": "MUSIC_SPOTIFY",
		},
		{
			"name": "Art",
			"code": "ART",
		},
		{
			"name": "Reading",
			"code": "READING",
		},
		{
			"name": "Gaming",
			"code": "GAMING",
		},
		{
			"name": "Technology",
			"code": "TECHNOLOGY",
		},
		{
			"name": "Fashion",
			"code": "FASHION",
		},
		{
			"name": "Cooking",
			"code": "COOKING",
		},
		{
			"name": "Sports",
			"code": "SPORTS",
		},
	}

	db.Exec("DELETE FROM master_interest;")
	tx := db.Begin()
	for _, value := range data {
		m := &model.MasterInterestModel{
			Code: value["code"],
			Name: value["name"],
		}
		tx.Create(m)
	}

	return tx.Commit().Error
}
