package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	crops := getAllCrops()
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":        "Home Page",
			"payload":      crops,
			"is_logged_in": loggedIn,
		},
	)

}

func getCrop(c *gin.Context) {
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)
	if cropID, err := strconv.Atoi(c.Param("crop_id")); err == nil {
		if crop, err := getCropByID(cropID); err == nil {
			c.HTML(
				http.StatusOK,
				"crop.html",
				gin.H{
					"title":        crop.Name,
					"payload":      crop,
					"is_logged_in": loggedIn,
				},
			)

		} else {
			// If the crop is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid crop ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}

func showCropCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Add New Crop"}, "create-crop.html")
}

func createCrop(c *gin.Context) {
	name := c.PostForm("name")
	start_date := c.PostForm("start_date")
	environment := c.PostForm("environment")
	qty := c.PostForm("qty")
	notes := c.PostForm("notes")

	if a, err := createNewCrop(name, start_date, environment, qty, notes); err == nil {
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
