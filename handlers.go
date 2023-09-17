package main

import (
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MetaHandler provives access to GET /meta end-point
func MetaHandler(c *gin.Context) {
	data := metadata("")
	c.AsciiJSON(http.StatusOK, data)
}

// MetaSiteHandler provides access to GET /meta/:site end-point
func MetaSiteHandler(c *gin.Context) {
	site := c.Param("site")
	data := metadata(site)
	c.AsciiJSON(http.StatusOK, data)
}

// MetaPostHandler provides access to POST /meta end-point
func MetaPostHandler(c *gin.Context) {
	var meta MetaData
	err := c.BindJSON(&meta)
	if err == nil {
		if meta.ID == "" {
			if uuid, err := uuid.NewRandom(); err == nil {
				meta.ID = hex.EncodeToString(uuid[:])
			}
		}
		_metaData = append(_metaData, meta)
		c.JSON(200, gin.H{"status": "ok"})
	} else {
		c.JSON(400, gin.H{"status": "fail", "error": err.Error()})
	}
}
