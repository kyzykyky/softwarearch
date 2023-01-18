package fiberpreset

import (
	"fmt"

	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kyzykyky/softwarearch/svcreg/pkg/logger"
	"go.uber.org/zap"
)

type Server struct {
	Service string
	Title   string
	Host    string
	Port    int
	App     *fiber.App
	Log     *zap.Logger
}

// Create sample fiber app without service routes
func (s *Server) New() error {
	s.App = fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  s.Service,
		AppName:       s.Title,
	})
	s.Log = logger.Logger().Named(fmt.Sprintf("%s.server", s.Service))

	s.App.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: s.RecoverStackTraceHandler,
	}))
	s.App.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	s.App.Use(fiberLogger.New(logconf))

	// Consul health check
	s.App.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "ok"})
	})
	return nil
}
