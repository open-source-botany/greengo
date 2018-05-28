package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	crops := getAllCrops()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": crops,
		},
	)

}

func getCrop(c *gin.Context) {
	if cropID, err := strconv.Atoi(c.Param("crop_id")); err == nil {
		if crop, err := getCropByID(cropID); err == nil {
			c.HTML(
				http.StatusOK,
				"crop.html",
				gin.H{
					"title":   crop.Name,
					"payload": crop,
				},
			)

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
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
