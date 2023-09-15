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
		data := metadata("")
		c.AsciiJSON(http.StatusOK, data)
	})
	r.GET("/meta/:site", func(c *gin.Context) {
		site := c.Param("site")
		data := metadata(site)
		c.AsciiJSON(http.StatusOK, data)
	})
	r.POST("/meta", func(c *gin.Context) {
		var meta MetaData
		err := c.BindJSON(&meta)
		if err == nil {
			_metaData = append(_metaData, meta)
			c.JSON(200, gin.H{"status": "ok"})
		} else {
			c.JSON(400, gin.H{"status": "fail", "error": err.Error()})
		}
	})

	return r
}

func Server(configFile string) {
	r := setupRouter()
	r.Run(":9092")
}
