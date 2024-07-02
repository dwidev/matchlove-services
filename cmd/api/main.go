package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"log"
	server "matchlove-services/internal"
	"matchlove-services/pkg/config"
	"matchlove-services/pkg/database"
)

func main() {
	migrate := flag.Bool("migrate", false, "migrate db")
	seeder := flag.Bool("seed", false, "seeding db data")
	flag.Parse()

	cfg := config.Load()
	db, err := database.Open(cfg.DB_DSN)
	if err != nil {
		logrus.Fatal("error open db : ", err)
	}

	dbs := db.Instance()
	srv := server.New(dbs)

	if ok := runMigration(db, *migrate); ok {
		return
	}

	if ok := runSeeder(db, *seeder); ok {
		return
	}

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server on port %s: %s", cfg.ServerPort, err)
	}
}

func runMigration(db *database.Database, migrate bool) bool {
	if migrate {
		if err := db.AutoMigrate(); err != nil {
			logrus.Fatal("error on migration database : ", err)
		}
		logrus.Info("migrate db success")
		return true
	}

	return false
}

func runSeeder(db *database.Database, seed bool) bool {
	if seed {
		if err := db.Seeder(); err != nil {
			logrus.Errorf("seed db error : %v", err)
			return true
		}

		logrus.Info("seed db success")
		return true
	}

	return false
}
