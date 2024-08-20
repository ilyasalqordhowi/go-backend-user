package routers

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.ListAllUsers)
	routerGroup.GET("/:id",controllers.DetailUsers)
	routerGroup.POST("/",controllers.CreateUsers)
	routerGroup.PATCH("/:id",controllers.UpdateUser)
	routerGroup.DELETE("/:id",controllers.DeleteUsers)
}