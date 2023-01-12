package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	// Path to the SQLite database file
	Path string
}

type Repository struct {
	DbConnection *gorm.DB
}

func (c Config) Connect() (Repository, error) {
	db, err := gorm.Open(sqlite.Open(c.Path), &gorm.Config{})
	if err != nil {
		return Repository{}, err
	}
	return Repository{db}, nil
}
