package main

func initializeRoutes() {

	router.GET("/", showIndexPage)
	router.GET("/crop/view/:crop_id", getCrop)
}
