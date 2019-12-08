package model

import (
	. "github.com/saichler/messaging/golang/net/protocol"
	. "github.com/saichler/utils/golang"
)

type Inventory struct {
	SID      *ServiceID
	Services []*ServiceID
}

func (inv *Inventory) ToBytes() []byte {
	bs := NewByteSlice()
	inv.Write(bs)
	return bs.Data()
}

func (inv *Inventory) FromBytes(data []byte) {
	bs := NewByteSliceWithData(data, 0)
	inv.Read(bs)
}

func (inv *Inventory) Write(bs *ByteSlice) {
	inv.SID.Write(bs)
	bs.AddInt(len(inv.Services))
	for _, sid := range inv.Services {
		sid.Write(bs)
	}
}

func (inv *Inventory) Read(bs *ByteSlice) {
	inv.SID = &ServiceID{}
	inv.SID.Read(bs)
	size := bs.GetInt()
	inv.Services = make([]*ServiceID, size)
	for i := 0; i < size; i++ {
		inv.Services[i] = &ServiceID{}
		inv.Services[i].Read(bs)
	}
}
