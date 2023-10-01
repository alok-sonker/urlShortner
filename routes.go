package main

import (
	"github.com/gin-gonic/gin"
)

func initRoutes(server *gin.Engine) {
	server.GET("/healthCheck", healthCheck)
	server.POST("/createURL", createURL)
	server.GET("/getURL", getURL)
	server.GET("/getAllURL", getAllURL)
	server.GET("/metrics", getMetrics)
	server.GET("/", indexHandler)
	server.POST("/", formHandler)
}
