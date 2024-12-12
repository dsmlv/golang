package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware ensures the user has the required role to access the endpoint
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the role from the context (set by AuthMiddleware)
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden: no role found"})
			c.Abort()
			return
		}

		// Check if the user's role matches the required role
		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden: insufficient privileges"})
			c.Abort()
			return
		}

		c.Next()
	}
}
