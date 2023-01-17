package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/kyzykyky/softwarearch/svcreg/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger = logger.Logger().Named("consul.connector")

func RegisterService(id, name string, tags []string, port int) error {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Error("consul: error creating client",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return err
	}

	registration := &api.AgentServiceRegistration{
		ID:   id,
		Name: name,
		Tags: tags,
		Port: port,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://localhost:8000/health",
			Interval: "5s",
		},
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		log.Error("consul: error registering service",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return err
	}

	services, err := client.Agent().Services()
	if err != nil {
		log.Error("consul: error getting services",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return err
	}
	log.Debug("consul: registered services",
		zapcore.Field{Key: "services", Type: zapcore.ReflectType, Interface: services})
	return nil
}
