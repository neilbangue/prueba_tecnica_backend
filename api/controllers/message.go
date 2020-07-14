package controllers

import (
	"server/models"

	"github.com/gin-gonic/gin"
)

func Messages(c *gin.Context) {
	var message models.Message
	c.Bind(&message)

	response := &models.Response{Count: len(message.Text)}
	c.JSON(200, response)
}
