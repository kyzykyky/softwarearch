package gorm_bookrepository

import (
	"github.com/kyzykyky/softwarearch/bookservice/internal/data/bookrepository"
	"gorm.io/gorm"
)

type Config struct {
	DbConnection *gorm.DB
}

type bookrepo struct {
	db *gorm.DB
}

func (c Config) Init() (bookrepository.BookRepository, error) {
	c.DbConnection.AutoMigrate(&Book{})
	return bookrepo{db: c.DbConnection}, nil
}
