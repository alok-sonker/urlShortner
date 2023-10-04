package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func initServer() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	return gin.Default()
}

func run(server *gin.Engine) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprintf("%v", PORT)
	}
	fmt.Println("Listening and serving HTTP on :" + port + " ... ")
	err := server.Run(":" + port + "")
	if err != nil {
		return err
	}
	return nil
}

func loadViews(server *gin.Engine) {
	server.LoadHTMLGlob("views/*")
}
