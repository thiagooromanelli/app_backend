package routes

import (
	"app.com/backend/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.HandleGetUsers)
			users.GET("/:id", controllers.HandleGetUserById)
			users.POST("/", controllers.HandlePostUsers)
			// users.PUT("/", controllers.)
			// users.DELETE("/", controllers.HandleGetUsers)
		}
	}

	return router
}
