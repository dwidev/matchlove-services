package main

import (
	"context"
	"flag"
	"github.com/sirupsen/logrus"
	"log"
	server "matchlove-services/internal"
	"matchlove-services/pkg/cache"
	"matchlove-services/pkg/config"
	"matchlove-services/pkg/database"
	"matchlove-services/pkg/database/seeder"
)

func main() {
	ctx := context.Background()
	// flag
	migrate := flag.Bool("migrate", false, "migrate db")
	s := flag.Bool("seed", false, "seeding db data")
	master := flag.Bool("master", false, "master address")
	user := flag.Bool("user", false, "user address")
	flag.Parse()

	// config
	cfg := config.Load()

	// database
	db, err := database.Open(cfg.DB_DSN)
	if err != nil {
		logrus.Fatal("error open db : ", err)
	}
	dbs := db.Instance()

	// cache
	c := cache.New(cache.RedisCache)
	defer func() {
		err := c.Disconnect(ctx)
		if err != nil {
			logrus.Fatal("error disconnect cache.RedisCache : ", err)
		}
	}()

	// migration
	if ok := runMigration(db, *migrate); ok {
		return
	}

	// seeder
	seederType := seeder.RunType{
		Seed:   *s,
		Master: *master,
		User:   *user,
	}
	if ok := runSeeder(db, seederType); ok {
		return
	}

	srv := server.New(dbs, c)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server on port %s: %s", cfg.ServerPort, err)
	}
}

func runMigration(db *database.Database, migrate bool) bool {
	if migrate {
		if err := db.Migration(); err != nil {
			logrus.Fatal("error on migration database : ", err)
		}
		logrus.Info("migrate db success")
		return true
	}

	return false
}

func runSeeder(db *database.Database, seederType seeder.RunType) bool {
	if seederType.Seed {
		if err := db.Seeder(seederType); err != nil {
			logrus.Errorf("seed db error : %v", err)
			return true
		}

		logrus.Info("seed db success")
		return true
	}

	return false
}
