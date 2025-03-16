package routes

import (
	"go-web-suite/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	{
		api.POST("/register", controllers.Register)
	}

	return router
}
