package server

import (
	"os"
	"os/signal"

	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/logger"
	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
	"go.uber.org/zap/zapcore"
)

type Controller interface {
	NewServer(svc domain.BookService, conf Conf) Server
	StartServer(Server)
}

type Server interface {
	Listen() error
	StartupTasks() error
	ShutdownTasks(chan os.Signal)
}

type Conf interface{}

func StartServer(app Server) {
	logger.Logger().Info("Starting Bookservice API")

	err := app.StartupTasks()
	if err != nil {
		logger.Logger().Panic(
			"Failed to run startup tasks for Bookservice API",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	}

	// Shutdown listener
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go app.ShutdownTasks(c)

	if err := app.Listen(); err != nil {
		logger.Logger().Panic(
			"Failed to start Bookservice API",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		)
	}
}
