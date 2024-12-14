package main

import (
	"final/config"
	"final/migrations"
	"final/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
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

	// CORS Middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},        // Frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // HTTP methods
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Apply CORS middleware to Gin
	handler := corsMiddleware.Handler(router)

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
