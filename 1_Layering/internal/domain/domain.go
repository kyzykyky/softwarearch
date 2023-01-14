package domain

import (
	"github.com/kyzykyky/softwarearch/bookservice/internal/data/bookrepository"
	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/mq"
)

// Bussiness logic layer
type IDomain interface {
	// Methods are defined here
}

type Domain struct {
	BookDAO bookrepository.BookRepository
	MQ      mq.MQ
}
