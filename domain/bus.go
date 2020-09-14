package domain

import "github.com/HETIC-MT-P2021/correction-cqrs/cqrs"

var CommandBus *cqrs.CommandBus
var QueryBus *cqrs.QueryBus

func InitBusses() {
	CommandBus = cqrs.NewCommandBus()
	QueryBus = &cqrs.QueryBus{}

	_ = CommandBus.RegisterHandler(NewCreateOrderCommandHandler(), &CreateOrderCommand{})
}
