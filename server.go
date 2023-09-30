package main

import "github.com/gin-gonic/gin"

func initServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}
