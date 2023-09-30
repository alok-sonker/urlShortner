package main

import (
	"fmt"

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
	err := server.Run(":" + fmt.Sprintf("%v", PORT) + "")
	if err != nil {
		return err
	}
	return nil
}
