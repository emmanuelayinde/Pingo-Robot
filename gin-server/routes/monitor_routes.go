package routes

import (
	"github.com/emmanuelayinde/pingo-robot/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/monitor", controllers.CreateMonitor)
		v1.GET("/monitor/:id", controllers.GetMonitor)
	}

	return r
}
