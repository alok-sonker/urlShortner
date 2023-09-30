package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func initRoutes(server *gin.Engine) {
	fmt.Println("Listening and serving HTTP on :8080 ... ")
	server.GET("/healthCheck", healthCheck)
	server.POST("/createURL", createURL)
	server.GET("/getURL", getURL)
	server.GET("/getAllURL", getAllURL)
}
