package consul

import (
	"sync"

	"github.com/hashicorp/consul/api"
	"github.com/kyzykyky/softwarearch/svcreg/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger = logger.Logger().Named("consul.connector")

var lock = &sync.Mutex{}
var Client *api.Client

type Consul struct {
	log         *zap.Logger
	ServiceId   string
	ServiceName string
	Host        string
	Port        int
	Tags        []string
}

func GetClient() (*api.Client, error) {
	if Client == nil {
		lock.Lock()
		defer lock.Unlock()
		if Client == nil {
			config := api.DefaultConfig()
			client, err := api.NewClient(config)
			if err != nil {
				log.Error("consul: error creating client",
					zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
				return nil, err
			}
			Client = client
		}
	}
	return Client, nil
}
