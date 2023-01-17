package fiberpreset

import (
	"os"
	"time"

	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kyzykyky/softwarearch/svcreg/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap/zapcore"
)

func (s Server) RecoverStackTraceHandler(c *fiber.Ctx, err interface{}) {
	logger.Logger().Error("fiber: server error",
		zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err},
		zapcore.Field{Key: "path", Type: zapcore.StringType, String: c.Path()})
}

var logconf = fiberLogger.Config{
	Next:         nil,
	Format:       "[${time}] ${status} - ${method} [${path}] [${locals:User}@${ip}:${port}] ${latency}\n",
	TimeFormat:   "02.01.2006 15:04:05.000",
	TimeZone:     "Local",
	TimeInterval: 500 * time.Millisecond,
	Output:       os.Stderr,
}
