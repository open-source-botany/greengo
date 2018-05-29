package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var router *gin.Engine

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	// Routes
	initializeRoutes()
	// Start app server
	router.Run(":" + port)

}
