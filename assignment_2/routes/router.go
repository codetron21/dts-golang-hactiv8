package routes

import (
	"assignment_2/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(ctl controller.Controller) error {
	r := gin.Default()

	r.GET("/orders", ctl.GetOrders)
	r.POST("/orders", ctl.CreateOrder)
	r.PUT("/orders/:orderId", ctl.UpdateOrderById)
	r.DELETE("/orders/:orderId", ctl.DeleteOrderById)

	return r.Run("localhost:8080")
}
