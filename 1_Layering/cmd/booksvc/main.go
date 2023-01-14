package main

import (
	"github.com/kyzykyky/softwarearch/bookservice/internal/controller/server"
	"github.com/kyzykyky/softwarearch/bookservice/internal/controller/server/fiber"
	gorm_bookrepository "github.com/kyzykyky/softwarearch/bookservice/internal/data/bookrepository/gorm"
	"github.com/kyzykyky/softwarearch/bookservice/internal/data/config/gorm/sqlite"
	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/logger"
	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/mq/nats"
	"github.com/kyzykyky/softwarearch/bookservice/internal/service"

	"go.uber.org/zap/zapcore"
)

func main() {
	sqliteconn, err := sqlite.Config{
		Path: "books.db",
	}.Connect()
	if err != nil {
		logger.Logger().Panic("Sqlite gorm connection failed",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
	}
	bookDAO, err := gorm_bookrepository.Config{DbConnection: sqliteconn.DbConnection}.Init()
	if err != nil {
		logger.Logger().Panic("Sqlite gorm BookDAO Initialization failed",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
	}
	MQ, err := nats.Connection{Host: "localhost:4222"}.NewMQ()
	if err != nil {
		logger.Logger().Panic("Nats connection failed",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
	}

	Service, err := service.NewService(service.Service{
		BookDAO: bookDAO,
		MQ:      MQ,
	})
	if err != nil {
		logger.Logger().Panic("Service start failed",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
	}

	fiberServer := fiber.NewServer(Service, fiber.Conf{
		Host: "localhost:3000",
	})
	server.StartServer(&fiberServer)
}
