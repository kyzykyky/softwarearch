package http

import (
	nethttp "net/http"
	"time"

	"github.com/kyzykyky/bookservice/pkg/domain"
)

type Server struct {
	Host string
}

func (server *Server) Start(service *domain.Service) error {
	httpServer := &nethttp.Server{
		Addr:           server.Host,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Add handlers

	return httpServer.ListenAndServe()
}
