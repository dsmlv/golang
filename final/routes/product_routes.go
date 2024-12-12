package routes

import (
	"final/controllers"
	"final/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine) {
	// Protect product routes with AuthMiddleware
	productGroup := router.Group("/products", middlewares.AuthMiddleware())
	{
		productGroup.GET("/", controllers.GetProducts)   // Authenticated users can view products
		productGroup.GET("/:id", controllers.GetProduct) // Authenticated users can view product details

		// Admin-only routes
		adminGroup := productGroup.Group("/")
		adminGroup.Use(middlewares.RoleMiddleware("admin")) // Restrict these routes to admin users
		{
			adminGroup.POST("/", controllers.CreateProduct)      // Admin can create a product
			adminGroup.PUT("/:id", controllers.UpdateProduct)    // Admin can update a product
			adminGroup.DELETE("/:id", controllers.DeleteProduct) // Admin can delete a product
		}
	}
}
