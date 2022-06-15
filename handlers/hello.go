package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	log.Println("Hello World")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
