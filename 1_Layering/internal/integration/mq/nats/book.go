package nats

import (
	"encoding/json"

	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
)

const bookcreated_subject = "BOOK.CREATED"

func (mq MQ) PublishBook(book domain.Book) error {
	bookJson, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return mq.nc.Publish(bookcreated_subject, bookJson)
}
