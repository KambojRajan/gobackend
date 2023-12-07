package routes

import (
	controllers "management/Controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/orderItemId", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:orderId", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItemId", controllers.UpdateOrderItem())

}
