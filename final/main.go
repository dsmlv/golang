package main

import (
	"final/config"
	"final/migrations"
	"final/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Run migrations
	migrations.RunMigrations()

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(router)
	routes.RegisterProductRoutes(router)

	// Start the server
	router.Run(":8080")
}
