package database

import (
	"matchlove-services/internal/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(dsn string) (*Database, error) {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
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

func (d Database) Instance() *gorm.DB {
	return d.db
}

func (d Database) AutoMigrate() error {
	if err := d.db.AutoMigrate(
		&model.UserAccount{},
		&model.UserProfile{},
		&model.UserPreference{},
	); err != nil {
		return err
	}

	return nil
}
