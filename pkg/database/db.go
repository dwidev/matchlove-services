package database

import (
	"github.com/sirupsen/logrus"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/database/seeder"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(dsn string) (*Database, error) {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDb, errDb := database.DB()
	if errDb != nil {
		panic(errDb)
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxIdleTime(5 * time.Minute)
	sqlDb.SetConnMaxLifetime(60 * time.Minute)

	return &Database{
		db: database,
	}, nil
}

type Database struct {
	db *gorm.DB
}

func (d *Database) Instance() *gorm.DB {
	return d.db
}

func (d *Database) Migration() error {
	if err := d.db.AutoMigrate(
		&model.UserAccount{},
		&model.UserProfile{},
		&model.UserPreference{},
		&model.UserInterest{},
		&model.MasterInterestModel{},
		&model.MasterLookingFor{},
	); err != nil {
		return err
	}

	return nil
}

func (d *Database) Seeder(t seeder.RunType) error {
	if t.Master {
		err := seeder.SeedMasterInterest(d.db)
		if err != nil {
			logrus.Errorf("error seeding master interest: %v", err)
			return err
		}

		err = seeder.SeedMasterLookingFor(d.db)
		if err != nil {
			logrus.Errorf("error seeding master looking for: %v", err)
			return err
		}
	}

	if t.User {
		err := seeder.SeedUsers(d.db)
		if err != nil {
			logrus.Errorf("error seeding users: %v", err)
			return err
		}
	}

	return nil
}
