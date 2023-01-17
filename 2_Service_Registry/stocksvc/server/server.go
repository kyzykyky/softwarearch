package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	fiberpreset "github.com/kyzykyky/softwarearch/svcreg/pkg/fiberPreset"
	"go.uber.org/zap"
)

type Server struct {
	Host string
	Port int
	app  *fiber.App
	log  *zap.Logger
}

func (s Server) Start() error {
	server := fiberpreset.Server{
		Service: "stocksvc",
		Title:   "Stock",
		Host:    s.Host,
		Port:    s.Port,
	}
	err := server.New()
	if err != nil {
		return err
	}
	s.app = server.App
	s.log = server.Log

	s.SetRoutes()
	return s.app.Listen(s.Host + ":" + fmt.Sprint(s.Port))
}
