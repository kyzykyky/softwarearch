package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	Host string
	app  *fiber.App
}

func (s Server) Start() error {
	s.app = fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Bookservice",
	})

	s.app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: RecoverStackTraceHandler,
	}))
	s.app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	s.app.Use(fiberLogger.New(GetLoggerConfig()))

	s.SetRoutes()

	return s.app.Listen(s.Host)
}
