package main

import (
	"fmt"
	"log"

	authz "github.com/OreCast/common/authz"
	oreConfig "github.com/OreCast/common/config"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// GET routes
	r.GET("/meta", MetaHandler)
	r.GET("/meta/:site", MetaSiteHandler)

	// all POST methods ahould be authorized
	authorized := r.Group("/")
	authorized.Use(authz.TokenMiddleware(oreConfig.Config.Authz.ClientId, oreConfig.Config.MetaData.Verbose))
	{
		authorized.POST("/meta", MetaPostHandler)
	}

	return r
}

func Server() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := setupRouter()
	sport := fmt.Sprintf(":%d", oreConfig.Config.MetaData.WebServer.Port)
	log.Printf("Start HTTP server %s", sport)
	r.Run(sport)
}
