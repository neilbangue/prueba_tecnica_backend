package mappings

import (
	"connectity/websocket/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.New()
	Router.Use(controllers.Cors())

	Router.GET("/ws", func(c *gin.Context) {
		controllers.Handler(c.Writer, c.Request)
	})
}
