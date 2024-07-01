package server

import (
	"fmt"
	"io"
	"matchlove-services/internal/router"
	"matchlove-services/pkg/config"
	"matchlove-services/pkg/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
)

type server struct {
	engine *fiber.App
	db     *gorm.DB
	config *config.Schema
}

func New(config *config.Schema, db *gorm.DB) *server {
	return &server{
		engine: fiber.New(),
		config: config,
		db:     db,
	}
}

func (s *server) Start() error {
	// setup
	s.setupMiddleware()
	s.setupRoutes()
	s.setupLogger()

	// healty check
	s.engine.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	// start the server
	logrus.Infof("HTTP server is listen with port %s", s.config.ServerPort)
	if err := s.engine.Listen(fmt.Sprintf(":%s", s.config.ServerPort)); err != nil {
		logrus.Fatalf("cannot running server with err %s", err)
		return err
	}

	return nil
}

func (s server) setupMiddleware() {
	revocer := middleware.RecoverPanicLogging()

	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type,Authorization",
		AllowCredentials: true,
	}))
	s.engine.Use(middleware.Logging)
	s.engine.Use(revocer)
}

func (s server) setupLogger() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10,
		MaxBackups: 1,
		MaxAge:     1,
		Compress:   true,
	}

	multiwritter := io.MultiWriter(os.Stdout, lumberjackLogger)
	logrus.SetOutput(multiwritter)

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetLevel(logrus.InfoLevel)
}

func (s server) setupRoutes() {
	router := &router.Router{
		Engine: s.engine,
		Config: s.config,
		DB:     s.db,
	}

	router.Build()
}
