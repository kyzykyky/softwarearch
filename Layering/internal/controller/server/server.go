package server

import "github.com/kyzykyky/bookservice/pkg/domain"

type Server interface {
	Start(*domain.Service) error
}
