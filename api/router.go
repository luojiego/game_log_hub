package api

import (
	"game_log_hub/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the API routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Basic health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Serve static files for the front-end
	router.Static("/assets", "./public")
	router.StaticFile("/", "./public/index.html")

	// API routes
	api := router.Group("/api")
	{
		// Login error routes
		loginErrors := api.Group("/login-errors")
		{
			loginErrors.POST("/", controllers.CreateLoginError)
			loginErrors.GET("/", controllers.GetLoginErrors)
			loginErrors.GET("/:id", controllers.GetLoginErrorByID)
			loginErrors.DELETE("/:id", controllers.DeleteLoginError)
		}
	}

	return router
}
