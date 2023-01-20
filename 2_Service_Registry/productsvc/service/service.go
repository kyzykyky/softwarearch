package service

import (
	"github.com/kyzykyky/softwarearch/svcreg/pkg/consul"
	"github.com/kyzykyky/softwarearch/svcreg/pkg/logger"
	"go.uber.org/zap"
)

type Service struct {
	log              *zap.Logger
	ConsulConnection consul.Consul
}

func (s *Service) Start() error {
	s.log = logger.Logger().Named("productsvc.service")
	return nil
}
