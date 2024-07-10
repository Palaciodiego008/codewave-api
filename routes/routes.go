package routes

import (
	"codewave/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	controllers.RegisterUserRoutes(router)

	// // Grupo de rutas protegidas
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
