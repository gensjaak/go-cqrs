package routes

import (
	"github.com/HETIC-MT-P2021/correction-cqrs/controllers"
	"github.com/gin-gonic/gin"
)

func addOrderRoutes(rg *gin.RouterGroup) {
	orders := rg.Group("/orders")

	//orders.GET("/", controllers.ListOrders)
	orders.GET("/", controllers.CreateOrder)
}
