package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	// Routes
	initializeRoutes()
	// Start app server
	router.Run()

}
