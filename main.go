package main

import (
	"github.com/gin-gonic/gin"
)

// Joke contains information about a single Joke
func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":8080")

}
