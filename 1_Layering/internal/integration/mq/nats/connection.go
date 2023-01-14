package nats

import (
	imq "github.com/kyzykyky/softwarearch/bookservice/internal/integration/mq"
	"github.com/nats-io/nats.go"
)

type Connection struct {
	Host string
}

func (c Connection) NewMQ() (MQ, error) {
	nc, err := nats.Connect(c.Host)
	if err != nil {
		return MQ{}, err
	}
	return MQ{nc}, nil
}

type MQ struct {
	nc *nats.Conn
}

func (mq MQ) Status() (string, error) {
	status := mq.nc.Status()
	if status != nats.CONNECTED {
		return status.String(), imq.ErrConnFailed
	}
	return status.String(), nil
}

func (mq MQ) Disconnect() error {
	mq.nc.Close()
	return nil
}
