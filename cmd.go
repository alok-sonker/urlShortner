package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func execute() error {
	server := initServer()
	initRoutes(server)
	if err := run(server); err != nil {
		return err
	}
	return nil
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
