package routes

import (
	controllers "management/Controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controllers.GetMenus())
	incomingRoutes.GET("/menus/menuId", controllers.GetMenu())
	incomingRoutes.POST("/menus", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:menuId", controllers.UpdateMenu())
}
