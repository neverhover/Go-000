package main

import (
	"github.com/gin-gonic/gin"
	"github.com/neverhover/Go-000/Week02/api"
	"github.com/neverhover/Go-000/Week02/storage"
)

func main() {
	defer func() {
		releaseResource()
	}()
	// Init storage instance
	initStorage()

	// Init web service
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	registerRoutes(r)
	
	r.Run(":8081")
}

func initStorage() {
	storage.InitDB()
}

func releaseResource() {
	storage.CloseDB()
}

func registerRoutes(r *gin.Engine)  {
	r.GET("/user/:id", api.GetUser())
}