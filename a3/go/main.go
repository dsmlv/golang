package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go" // Import the JWT package for token handling
	"github.com/gin-contrib/cors" // Import CORS middleware for handling cross-origin requests
	"github.com/gin-gonic/gin"    // Import the Gin web framework
)

// Secret key used for signing JWT tokens
var jwtKey = []byte("secret")

// In-memory data store to simulate a database
var data = make(map[string]string)

// Dosmailova Dinara
// User model to define the structure of a user, including roles for RBAC
type User struct {
	Username string
	Password string
	Role     string // Role field for role-based access control
}

// Dosmailova Dinara
// Users based on the role(admin or simple user)
var users = map[string]User{
	"admin": {Username: "admin", Password: "password", Role: "admin"},
	"user":  {Username: "user", Password: "password", Role: "user"},
}

// Struct to hold user credentials during login
type Credentials struct {
	Username string `json:"username"` // JSON mapping for request payload
	Password string `json:"password"` // JSON mapping for request payload
}

// Struct for JWT claims, including username, role, and standard claims like expiration time
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// Function to generate a JWT token for a user
func generateJWT(username, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour) // Set token expiration to 1 hour
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Set expiration time in Unix format
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Create a new JWT token
	return token.SignedString(jwtKey)                          // Sign the token with the secret key
}

// Handler for user login and token generation
func login(c *gin.Context) {
	var creds Credentials

	// Bind the incoming JSON to the credentials struct and check for errors
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the user exists and the password is correct
	user, exists := users[creds.Username]
	if !exists || user.Password != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token for the authenticated user
	token, err := generateJWT(user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the generated token in the response
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Middleware to protect routes and validate JWT tokens
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Parse the token and extract claims
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil // Provide the secret key for validation
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the username and role in the context for use in handlers
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next() // Proceed to the next handler
	}
}

// Dosmailova Dinara
// Middleware for role-based authorization
func roleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the role from the context
		role := c.MustGet("role").(string)
		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
			c.Abort()
			return
		}
		c.Next() // Proceed to the next handler
	}
}

func main() {
	r := gin.Default() // Create a new Gin router

	// Enable CORS with default settings
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Your React app origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public route: User login
	r.POST("/login", login)

	// Protected routes group
	protected := r.Group("/")
	protected.Use(authMiddleware()) // Apply the auth middleware to protect routes

	// Endpoint to create a new item in the data store
	protected.POST("/create", func(c *gin.Context) {
		var request struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data[request.ID] = request.Name // Add item to the data store
		c.JSON(http.StatusOK, gin.H{"message": "Created successfully"})
	})

	// Endpoint to read an item from the data store by ID
	protected.GET("/read/:id", func(c *gin.Context) {
		id := c.Param("id")
		if name, exists := data[id]; exists {
			c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		}
	})

	// Endpoint to update an item in the data store
	protected.PUT("/update/:id", func(c *gin.Context) {
		id := c.Param("id")
		var request struct {
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if _, exists := data[id]; exists {
			data[id] = request.Name // Update the item in the data store
			c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		}
	})

	// Endpoint to delete an item from the data store
	protected.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		if _, exists := data[id]; exists {
			delete(data, id) // Remove the item from the data store
			c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		}
	})

	// Admin-only route, protected by role-based access control
	protected.GET("/admin", roleMiddleware("admin"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin!"})
	})

	// Start the server on port 8080
	r.Run(":8080")
}
