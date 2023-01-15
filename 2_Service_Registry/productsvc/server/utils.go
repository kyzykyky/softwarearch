package server

import (
	"io"
	"os"
	"time"

	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kyzykyky/softwarearch/svcreg/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger = logger.Logger().Named("productsvc.server")

func RecoverStackTraceHandler(c *fiber.Ctx, err interface{}) {
	logger.Logger().Error("fiber: server error",
		zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		zapcore.Field{Key: "path", Type: zapcore.StringType, String: c.Path()})
}

var accessFile *os.File

var conf = fiberLogger.Config{
	Next:         nil,
	Format:       "[${time}] ${status} - ${method} [${path}] [${locals:User}@${ip}:${port}] ${latency}\n",
	TimeFormat:   "02.01.2006 15:04:05.000",
	TimeZone:     "Local",
	TimeInterval: 500 * time.Millisecond,
}

func GetLoggerConfig() fiberLogger.Config {
	var err error
	accessFile, err = os.OpenFile("log/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0764)
	if err != nil {
		log.Error("fiber: error opening file for logger",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		conf.Output = os.Stderr
		return conf
	}
	conf.Output = io.MultiWriter(os.Stderr, accessFile)
	return conf
}
