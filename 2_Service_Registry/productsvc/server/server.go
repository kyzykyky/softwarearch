package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kyzykyky/softwarearch/svcreg/pkg/consul"
	fiberpreset "github.com/kyzykyky/softwarearch/svcreg/pkg/fiberPreset"
	"github.com/kyzykyky/softwarearch/svcreg/productsvc/service"
	"go.uber.org/zap"
)

type Server struct {
	ServiceId        string
	Host             string
	Port             int
	ConsulConnection consul.Consul
	app              *fiber.App
	log              *zap.Logger

	Service service.Service
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

	s.Service.Start()
	s.ConsulConnection, err = consul.RegisterService(s.ServiceId, title, s.Host, s.Port, tags)
	if err != nil {
		return err
	}
	s.Service.ConsulConnection = s.ConsulConnection

	s.SetRoutes()

	s.log.Info(fmt.Sprintf("Service %s starting", s.ServiceId))
	return s.app.Listen(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
