package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joelgarciajr84/go-rest-api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		heroes := main.Group("heroes")
		{
			heroes.GET("/", controllers.ShowBestHero)
		}
	}
	return router
}
