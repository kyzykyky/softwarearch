package mq

type MQ interface {
	Publish(string, []byte) error
	Subscribe(string, func([]byte)) error
}
