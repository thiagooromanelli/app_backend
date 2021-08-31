package routes

import (
	"app.com/backend/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.GetOk)
		}
	}

	return router
}
