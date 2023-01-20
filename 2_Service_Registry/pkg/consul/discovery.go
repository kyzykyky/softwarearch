package consul

import (
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap/zapcore"
)

// DiscoverService returns a list of services with the given name
func (con *Consul) DiscoverService(service string) ([]*api.AgentService, error) {
	services, err := Client.Agent().Services()
	if err != nil {
		con.log.Error("consul: error getting services",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return nil, err
	}
	var serviceList []*api.AgentService
	for _, s := range services {
		if s.Service == service {
			serviceList = append(serviceList, s)
		}
	}
	return serviceList, nil
}
