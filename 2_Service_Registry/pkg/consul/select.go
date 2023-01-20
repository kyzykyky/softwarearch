package consul

import (
	"fmt"
	"sync"
)

var nextIndexMap map[string]int = make(map[string]int)
var mutex = &sync.Mutex{}

// Get the next service instance by name with round-robin selection
func (con *Consul) SelectServiceInstance(service string) (string, error) {
	services, err := con.DiscoverService(service)
	if err != nil {
		return "", err
	}
	mutex.Lock()
	defer mutex.Unlock()
	nextIndex, ok := nextIndexMap[service]
	if !ok || nextIndex >= len(services) {
		nextIndexMap[service] = 0
	}
	instance := services[nextIndexMap[service]]
	nextIndexMap[service]++

	address := fmt.Sprintf("%s:%d", instance.Address, instance.Port)
	return address, nil
}
