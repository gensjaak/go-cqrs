package domain

import (
	"errors"
	"fmt"
	"github.com/HETIC-MT-P2021/correction-cqrs/cqrs"
)

type CreateOrderCommand struct {
	Price uint
}

func (coc CreateOrderCommand) Payload() interface{} {
	return ""
}

func (coc CreateOrderCommand) CommandType() string {
	return "CreateOrderCommand"
}

type CreateOrderCommandHandler struct {}

func (ch CreateOrderCommandHandler) Handle(command cqrs.CommandMessage) error {

	switch cmd := command.Payload().(type) {
		case *CreateOrderCommand:
			fmt.Println("Handler.")
			fmt.Println(cmd.Price)
	default:
		return errors.New("Bad commadn type")
	}

	return nil
}

func NewCreateOrderCommandHandler() *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{}
}