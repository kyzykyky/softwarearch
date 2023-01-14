package service

import (
	"os"
	"os/signal"

	"github.com/kyzykyky/softwarearch/bookservice/internal/data/bookrepository"
	"github.com/kyzykyky/softwarearch/bookservice/internal/domain"
	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/logger"
	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/mq"
	"go.uber.org/zap/zapcore"
)

// Initial configuration
type Service struct {
	BookDAO bookrepository.BookRepository
	MQ      mq.MQ
}

// Logic owner
type service struct {
	Domain  domain.IDomain
	BookDAO bookrepository.BookRepository
	MQ      mq.MQ
}

func NewService(serv Service) (service, error) {
	service := service{
		BookDAO: serv.BookDAO,
		MQ:      serv.MQ,
		Domain: domain.Domain{
			BookDAO: serv.BookDAO,
		},
	}
	err := service.start()
	if err != nil {
		logger.Logger().Error("Service start failed")
		return service, err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go service.stop(c)

	return service, nil
}

// Start tasks
func (s service) start() error {
	mqstatus, err := s.MQ.Status()
	if err != nil {
		logger.Logger().Error("MQ status failed",
			zapcore.Field{Key: "status", Type: zapcore.StringType, String: mqstatus},
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return err
	}
	return nil
}

// Stop tasks
func (s service) stop(c chan os.Signal) {
	<-c
	err := s.MQ.Disconnect()
	if err != nil {
		logger.Logger().Error("MQ disconnect failed",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
	}
}
