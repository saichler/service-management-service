package commands

import (
	. "github.com/saichler/console/golang/console/commands"
	. "github.com/saichler/service-manager/golang/service-manager"
	. "github.com/saichler/utils/golang"
	"strings"
)

type Shutdown struct {
	service IService
}

func NewShutdown(sm IService) *Shutdown {
	sd := &Shutdown{}
	sd.service = sm
	return sd
}

func (c *Shutdown) Command() string {
	return "shutdown"
}

func (c *Shutdown) Description() string {
	return "Shutdown the Service Manager"
}
func (c *Shutdown) Usage() string {
	return "shutdown"
}
func (c *Shutdown) ConsoleId() *ConsoleId {
	return c.service.ConsoleId()
}
func (c *Shutdown) RunCommand(args []string, id *ConsoleId) (string, *ConsoleId) {
	Print("Are you sure you want to shutdown " + c.service.ServiceManager().ConsoleId().ID() + " (yes/no)?")
	reply, _ := Read()
	reply = strings.ToLower(reply)
	for reply != "no" && reply != "yes" {
		Print("yes/no please?")
		reply, _ := Read()
		reply = strings.ToLower(reply)
	}
	if reply == "yes" {
		c.service.ServiceManager().Publish("Shutdown", c.service, []byte("Shutdown"))
		defer c.service.ServiceManager().Shutdown()
		return "Shutting Down...", nil
	}
	return "Canceled Shutdown.", nil
}
