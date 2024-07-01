package main

import (
	"github.com/sirupsen/logrus"
	"log"
	server "matchlove-services/internal"
	"matchlove-services/pkg/config"
	"matchlove-services/pkg/database"
)

type Test struct {
	Name string
}

func main() {
	cfg := config.Load()
	db, err := database.Open(cfg.DB_DSN)
	if err != nil {
		logrus.Fatal("error open db : ", err)
	}

	dbs := db.Instance()
	srv := server.New(cfg, dbs)

	if err := db.AutoMigrate(); err != nil {
		logrus.Fatal("error on mingration database : ", err)
	}

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server on port %s: %s", cfg.ServerPort, err)
	}
}
