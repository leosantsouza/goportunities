package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leosantsouza/goportunities/handler"
)

// initializeRoutes sets up the API routes for the goportunities application.
// It takes a pointer to a gin.Engine instance as a parameter, which represents the
// HTTP router. The function registers routes for different endpoints related to
// openings, such as creating, retrieving, updating, and deleting openings.
//
// Parameters:
// - router: A pointer to a gin.Engine instance representing the HTTP router.
func initializeRoutes(router *gin.Engine) {
    v1 := router.Group("/api/v1")
    {
        // Openings routes
        v1.GET("/opening", handler.ShowOpening)
        v1.POST("/opening", handler.CreateOpening)
        v1.DELETE("/opening", handler.DeleteOpening)
        v1.PUT("/opening", handler.UpdateOpening)
        v1.GET("/openings", handler.ListOpenings)
    }
}
	