package main

func initializeRoutes() {
	cropRoutes := router.Group("/crop")
	{
		cropRoutes.GET("/view/:crop_id", getCrop)

		cropRoutes.GET("/create", showCropCreationPage)

		cropRoutes.POST("/create", createCrop)
	}
	router.GET("/", showIndexPage)
}
