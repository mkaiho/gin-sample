package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_PORT = "3000"
)

func main() {
	port := os.Getenv("API_REQUEST_PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	engine := gin.Default()
	engine.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	engine.Run(":" + port)
}
