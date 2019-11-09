package commands

import (
	. "github.com/saichler/console/golang/console/commands"
	. "github.com/saichler/service-manager/golang/service-manager"
)

type LoadService struct {
	service IService
}

func NewLoadService(sm IService) *LoadService {
	sd := &LoadService{}
	sd.service = sm
	return sd
}
func (c *LoadService) Command() string {
	return "load-service"
}
func (c *LoadService) Description() string {
	return "Loads a service"
}
func (c *LoadService) Usage() string {
	return "load-service <path to plugin file>"
}
func (c *LoadService) ConsoleId() *ConsoleId {
	return c.service.ConsoleId()
}
func (c *LoadService) RunCommand(args []string, id *ConsoleId) (string, *ConsoleId) {
	c.service.ServiceManager().LoadService(args[0])
	return "", nil
}
