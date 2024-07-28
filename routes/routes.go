package routes

import (
	"codewave/controllers"
	"codewave/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply the CORS middleware
	router.Use(cors.New(config))

	// Apply the auth middleware to protected routes
	protected := router.Group("/")
	controllers.UserRoutes(protected)
	protected.Use(middleware.AuthRequired())
	{
		controllers.ProjectRoutes(protected)
		controllers.GeminiRoutes(protected)
		controllers.OpenAIRoutes(protected)
		controllers.AnalysisRoutes(protected)
	}

	return router
}
