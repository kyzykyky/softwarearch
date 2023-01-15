package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var lock = &sync.Mutex{}

var singleInstanceLogger *zap.Logger

func Logger() *zap.Logger {
	if singleInstanceLogger == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstanceLogger == nil {
			singleInstanceLogger = getZapLogger()
		}
	}
	return singleInstanceLogger
}

func getZapLogger() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("error building logger: %v\n", err)
	}
	return logger
}
