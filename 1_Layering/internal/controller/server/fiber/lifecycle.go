package fiber

import (
	"os"

	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/logger"
	"go.uber.org/zap/zapcore"
)

func (app *Server) StartupTasks() error {
	logger.Logger().Info("Starting fiber Bookservice API")
	// Any startup tasks here
	return nil
}

func (app *Server) ShutdownTasks(c chan os.Signal) {
	<-c
	logger.Logger().Info("Shutting down fiber Bookservice API")

	err := app.App.Shutdown()
	if err != nil {
		logger.Logger().Error("error shutting down fiber Bookservice API",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
	}

	accessFile.Close()
}
