package main

import (
	"fmt"
	"log"

	authz "github.com/OreCast/common/authz"
	oreConfig "github.com/OreCast/common/config"
	oreMongo "github.com/OreCast/common/mongo"
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
		authorized.DELETE("/meta/:mid", MetaDeleteHandler)
	}

	return r
}

func Server() {
	// init MongoDB
	oreMongo.InitMongoDB(oreConfig.Config.MetaData.DBUri)

	// setup web router and start the service
	r := setupRouter()
	sport := fmt.Sprintf(":%d", oreConfig.Config.MetaData.WebServer.Port)
	log.Printf("Start HTTP server %s", sport)
	r.Run(sport)
}
