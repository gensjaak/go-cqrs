package domain

import (
	"errors"
	"fmt"
	"github.com/HETIC-MT-P2021/correction-cqrs/cqrs"
)

type CreateOrderCommand struct {
	Price uint
}

type CreateOrderCommandHandler struct {}

func (ch CreateOrderCommandHandler) Handle(command cqrs.CommandMessage) error {

	switch cmd := command.Payload().(type) {
		case *CreateOrderCommand:
			fmt.Println("Handler.")
			fmt.Println(cmd.Price)
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewCreateOrderCommandHandler() *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{}
}
