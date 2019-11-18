package message_handlers

import (
	"github.com/saichler/messaging/golang/net/protocol"
	model2 "github.com/saichler/service-management-service/golang/management-service/model"
	management_service "github.com/saichler/service-management-service/golang/management-service/service"
	. "github.com/saichler/service-manager/golang/service-manager"
	utils "github.com/saichler/utils/golang"
)

type PingMH struct {
	ms   *management_service.ManagementService
	ping *protocol.Message
	//hash string
}

func NewPingMH(service IService) *PingMH {
	ping := &PingMH{}
	ping.ms = service.(*management_service.ManagementService)
	return ping
}

func (m *PingMH) Init() {
	m.ms.ServiceManager().ScheduleMessage(m, 10, 0)
}

func (m *PingMH) Topic() string {
	return "Ping"
}

func (m *PingMH) Message(destination *protocol.ServiceID, data []byte, isReply bool) *protocol.Message {
	dest := protocol.NewServiceID(protocol.NetConfig.PublishID(), m.ms.Topic(), m.ms.ID())
	return m.ms.ServiceManager().NewMessage(m.Topic(), m.ms, dest, m.inventory(), false)
}

func (m *PingMH) Handle(message *protocol.Message) {
	inv := &model2.Inventory{}
	bs := utils.NewByteSliceWithData(message.Data(), 0)
	inv.Object(bs)
	/*
		if len(inv.Services) > 0 {
			utils.Info("Reveived Inventory From:", message.Source().String(), " with:")
			for _, s := range inv.Services {
				utils.Info("  ", s.String())
			}
		}*/
	m.ms.ServiceManager().ServiceNetwork().UpdateInventory(inv)
}

func (m *PingMH) inventory() []byte {
	inv := &model2.Inventory{}
	inv.SID = m.ms.ServiceID()
	services := m.ms.ServiceManager().Services()
	inv.Services = make([]*protocol.ServiceID, 0)
	for _, service := range services {
		inv.Services = append(inv.Services, service.ServiceID())
	}
	bs := utils.NewByteSlice()
	inv.Bytes(bs)
	return bs.Data()
}

func (m *PingMH) Request(args ...interface{}) (interface{}, error) {
	return nil, nil
}
