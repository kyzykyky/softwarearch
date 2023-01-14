package config

type DBConnect interface {
	Connect() error
	Close() error
}
