package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kyzykyky/softwarearch/svcreg/pkg/consul"
	fiberpreset "github.com/kyzykyky/softwarearch/svcreg/pkg/fiberPreset"
	"go.uber.org/zap"
)

type Server struct {
	ServiceId string
	Host      string
	Port      int
	app       *fiber.App
	log       *zap.Logger
}

func (s Server) Start() error {
	service := "productsvc"
	title := "product"
	// urlprefix-<service>/ is used by Fabio to route requests
	tags := []string{"lab", title, fmt.Sprintf("urlprefix-%s/", service)}
	server := fiberpreset.Server{
		Service: s.ServiceId,
		Title:   title,
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

	err = consul.RegisterService(s.ServiceId, title, tags, s.Port)
	if err != nil {
		return err
	}
	s.log.Info(fmt.Sprintf("Service %s starting", s.ServiceId))
	return s.app.Listen(s.Host + ":" + fmt.Sprint(s.Port))
}
