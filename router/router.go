package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Initialize Gin and configure routes
	router := gin.Default()

	// Define routes
	initializeRoutes(router)

	// Run the server
	router.Run(":3000")

}
