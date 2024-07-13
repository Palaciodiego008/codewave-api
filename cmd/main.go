package main

import (
	"codewave/config"
	"codewave/routes"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	config.InitDB()
	router := routes.InitRoutes()
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
