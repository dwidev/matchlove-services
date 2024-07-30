package server

import (
	"fmt"
	"io"
	"matchlove-services/internal/router"
	"matchlove-services/pkg/cache"
	"matchlove-services/pkg/config"
	"matchlove-services/pkg/injection"
	"matchlove-services/pkg/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
)

type Server struct {
	engine *fiber.App
	db     *gorm.DB
	config *config.Schema
	cache  cache.Cache
}

func New(db *gorm.DB, c cache.Cache) *Server {
	return &Server{
		engine: fiber.New(),
		config: config.Get(),
		db:     db,
		cache:  c,
	}
}

func (s Server) Start() error {
	// setup
	s.setupMiddleware()
	s.setupRoutes()
	s.setupLogger()

	// health check
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

func (s Server) setupMiddleware() {
	panicLogging := middleware.RecoverPanicLogging()

	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type,Authorization",
		AllowCredentials: true,
	}))
	s.engine.Use(middleware.Logging)
	s.engine.Use(panicLogging)
}

func (s Server) setupLogger() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10,
		MaxBackups: 1,
		MaxAge:     1,
		Compress:   true,
	}

	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)
	logrus.SetOutput(multiWriter)

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetLevel(logrus.InfoLevel)
}

func (s Server) setupRoutes() {
	handler := injection.InitializeHandler(s.db, s.cache)
	route := &router.Router{
		Engine:  s.engine,
		Config:  s.config,
		Handler: handler,
		Cache:   s.cache,
	}

	route.Build()
}
