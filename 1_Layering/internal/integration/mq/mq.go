package mq

import (
	"errors"

	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
)

type Connection interface {
	NewMQ() (MQ, error)
}

type MQ interface {
	Status() (string, error)
	Disconnect() error

	PublishBook(domain.Book) error
}

var (
	ErrConnFailed error = errors.New("Connection failed")
)
