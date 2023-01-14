package logger

import (
	"io"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var lock = &sync.Mutex{}

var singleInstanceLogger *zap.Logger

func Logger() *zap.Logger {
	if singleInstanceLogger == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstanceLogger == nil {
			singleInstanceLogger = getZapLogger("cfg/logger.yaml")
		}
	}
	return singleInstanceLogger
}

func getZapLogger(configPath string) *zap.Logger {
	var usingDefaultConfig bool
	confFile, err := os.Open(configPath)
	if err != nil {
		log.Printf("error opening config file: %v\n", err)
		log.Println("Using default zap configuration")
		usingDefaultConfig = true
	}
	defer confFile.Close()
	conf, err := io.ReadAll(confFile)
	if err != nil && !usingDefaultConfig {
		log.Fatalf("error reading config file: %v\n", err)
	}

	var cfg zap.Config
	if !usingDefaultConfig {
		if err := yaml.Unmarshal(conf, &cfg); err != nil {
			log.Fatal(err)
		}
		err = os.MkdirAll("log", 0764)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("error building logger: %v\n", err)
	}
	return logger
}
