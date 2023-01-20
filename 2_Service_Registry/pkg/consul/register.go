package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap/zapcore"
)

func RegisterService(id, name, host string, port int, tags []string) (Consul, error) {
	client, err := GetClient()
	if err != nil {
		return Consul{}, err
	}

	registration := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Address: host,
		Port:    port,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health", host, port),
			Interval: "5s",
		},
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		log.Error("consul: error registering service",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return Consul{}, err
	}

	return Consul{
		log:         log.Named(id),
		ServiceId:   id,
		ServiceName: name,
		Host:        host,
		Port:        port,
		Tags:        tags,
	}, nil
}
