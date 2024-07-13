package routes

import (
	"codewave/controllers"
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

	controllers.HomeRoutes(router)
	controllers.UserRoutes(router)

	// auth := router.Group("/")
	// auth.Use(middleware.AuthMiddleware())
	// {
	// 	// Registra las rutas que requieren autenticación aquí
	// 	controllers.RegisterProjectRoutes(auth)
	// 	controllers.RegisterRecommendationRoutes(auth)
	// 	controllers.RegisterAnalysisRoutes(auth)
	// }

	return router
}
