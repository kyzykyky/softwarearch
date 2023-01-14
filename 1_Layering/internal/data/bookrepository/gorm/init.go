package gorm_bookrepository

import (
	"github.com/kyzykyky/softwarearch/bookservice/internal/data/bookrepository"
	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/logger"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type Config struct {
	DbConnection *gorm.DB
}

type bookrepo struct {
	db *gorm.DB
}

func (c Config) Init() (bookrepository.BookRepository, error) {
	err := c.DbConnection.AutoMigrate(&Book{})
	if err != nil {
		logger.Logger().Error("Failed to migrate Book table",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return bookrepo{}, err
	}
	return bookrepo{db: c.DbConnection}, nil
}
