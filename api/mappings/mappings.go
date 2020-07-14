package mappings

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(controllers.Cors())

	v1 := Router.Group("/v1")
	{
		v1.POST("/messages/", controllers.Messages)
	}
}
