package main

import (
	"log"
	server "matchlove-services/internal"
	"matchlove-services/pkg/config"
	"matchlove-services/pkg/database"

	"github.com/sirupsen/logrus"
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

	srv := server.New(cfg, db.Instance())

	if err := db.AutoMigrate(); err != nil {
		logrus.Fatal("error on mingration database : ", err)
	}

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server on port %s: %s", cfg.ServerPort, err)
	}
}
