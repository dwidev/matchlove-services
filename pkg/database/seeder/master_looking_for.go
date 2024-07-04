package seeder

import (
	"gorm.io/gorm"
	"matchlove-services/internal/model"
)

func SeedMasterLookingFor(db *gorm.DB) error {
	data := []map[string]string{
		{"code": "LT_PARTNER", "name": "A long-term partner 🥰💘"},
		{"code": "LOOKING_FRIENDS", "name": "Looking for friends 👋🏻🤙🏻"},
		{
			"code": "LOOKING_SIBLING",
			"name": "Looking for a brother or sister🙋🏻‍🙋🏻‍",
		},
		{"code": "FIGURING_IT_OUT", "name": "Still figuring it out 🤔"},
	}

	db.Exec("DELETE FROM master_looking_for;")

	tx := db.Begin()
	for _, value := range data {
		m := &model.MasterLookingFor{
			Code: value["code"],
			Name: value["name"],
		}
		tx.Create(m)
	}

	return tx.Commit().Error
}
