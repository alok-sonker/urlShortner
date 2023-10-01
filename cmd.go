package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func execute() error {
	server := initServer()
	loadViews(server)
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
	// server.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"GET", "HEAD", "OPTIONS", "POST", "PUT"},
	// 	AllowHeaders: []string{"Origin", "Access-Control-Allow-Headers", "Accept",
	// 		"X-Requested-With", "Content-Type", "Access-Control-Allow-Origin: *",
	// 		"Access-Control-Request-Method:*", "Access-Control-Request-Headers",
	// 		"Access-Control-Allow-Origin",
	// 		"Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	MaxAge:           12 * time.Hour,
	// 	AllowCredentials: true,
	// }))

	err := server.Run(":" + port + "")
	if err != nil {
		return err
	}
	return nil
}

func loadViews(server *gin.Engine) {
	server.LoadHTMLGlob("views/*")
}
