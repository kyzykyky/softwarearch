package fiber

import (
	"os"

	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/logger"
	"go.uber.org/zap/zapcore"
)

func (app *Server) StartupTasks() error {
	logger.Logger().Info("fiber: Starting Bookservice API")
	// Any startup tasks here
	return nil
}

func (app *Server) ShutdownTasks(c chan os.Signal) {
	<-c
	logger.Logger().Info("fiber: Shutting down Bookservice API")

	err := app.App.Shutdown()
	if err != nil {
		logger.Logger().Error("fiber: error shutting down Bookservice API",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
	}

	err = accessFile.Close()
	if err != nil {
		logger.Logger().Error("fiber: error closing access log file",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
	}
}
