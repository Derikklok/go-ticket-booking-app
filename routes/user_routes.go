package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Derikklok/go-ticket-booking-app/controllers"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/api/users")
	{
		user.GET("/", controllers.GetUsers)
		user.POST("/", controllers.CreateUser)
	}
}
