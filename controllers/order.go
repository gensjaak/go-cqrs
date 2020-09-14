package controllers

import (
	"github.com/HETIC-MT-P2021/correction-cqrs/cqrs"
	"github.com/HETIC-MT-P2021/correction-cqrs/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOrders(context *gin.Context) {
	context.JSON(http.StatusOK, "orders")
}

func CreateOrder(context *gin.Context) {
	command := cqrs.NewCommandMessage(&domain.CreateOrderCommand{Price: uint(10)})
	_ = domain.CommandBus.Dispatch(command)

	context.JSON(http.StatusOK, "orders")
}
