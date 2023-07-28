package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/meta", func(c *gin.Context) {
		data := metadata()
		c.AsciiJSON(http.StatusOK, data)
	})

	return r
}

func Server(configFile string) {
	r := setupRouter()
	r.Run(":9092")
}
