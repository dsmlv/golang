package main

import (
	"encoding/json" // For JSON encoding/decoding
	"fmt"           // For printing to console
	"log"           // For logging errors
	"net/http"      // For HTTP server functionality

	"github.com/gorilla/mux"  // Router for handling HTTP routes
	"github.com/rs/cors"      // CORS middleware to handle cross-origin requests
	"gorm.io/driver/postgres" // PostgreSQL driver for GORM
	"gorm.io/gorm"            // GORM ORM library
)

// Dosmailova Dinara
// Task struct represents a task entity in the database
type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"` // Primary key in the database
	Title       string `json:"title"`                // Task title
	Description string `json:"description"`          // Task description
	Completed   bool   `json:"completed"`            // Task completion status
}

var db *gorm.DB // Global variable to hold the database connection

// Dosmailova Dinara
// initDB initializes the PostgreSQL database connection
func initDB() {
	// Connection string for PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=taskmanager port=5432 sslmode=disable"
	var err error
	// Open the database connection with GORM and the PostgreSQL driver
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//Dosmailova Dinara
	// Automatically migrate the Task schema to create the tasks table
	db.AutoMigrate(&Task{})
}

// Dosmailova Dinara
// Handler to retrieve all tasks from the database
func getTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	// Find all tasks in the database
	db.Find(&tasks)
	// Respond with the tasks as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Dosmailova Dinara
// Handler to create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	// Decode the JSON request body into a Task struct
	json.NewDecoder(r.Body).Decode(&task)
	// Insert the new task into the database
	db.Create(&task)
	// Respond with the newly created task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Dosmailova Dinara
// Handler to update an existing task
func updateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	// Get task ID from the URL
	params := mux.Vars(r)
	// Find the task by ID
	db.First(&task, params["id"])
	// Decode the new data and update the task
	json.NewDecoder(r.Body).Decode(&task)
	db.Save(&task)
	// Respond with the updated task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Dosmailova Dinara
// Handler to delete a task by ID
func deleteTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	// Get task ID from the URL
	params := mux.Vars(r)
	// Delete the task by ID
	db.Delete(&task, params["id"])
	w.WriteHeader(http.StatusNoContent) // 204 No Content response
}

// Dosmailova Dinara
// Setup the routes for the API and start the server
func handleRequests() {
	// Create a new router instance
	router := mux.NewRouter()

	// Route to get all tasks
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	// Route to create a new task
	router.HandleFunc("/tasks", createTask).Methods("POST")
	// Route to update an existing task by ID
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	// Route to delete a task by ID
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	// Enable CORS to allow requests from the React frontend running on localhost:3000
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow requests from React app
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // Allow the required HTTP methods
		AllowedHeaders:   []string{"Content-Type"},                 // Allow JSON content-type
	})

	// Wrap the router with the CORS middleware
	handler := c.Handler(router)

	// Start the HTTP server on port 8080
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// Dosmailova Dinara
// Main function to initialize the database and start the server
func main() {
	initDB() // Initialize the database connection
	fmt.Println("Server running on port 8080")
	handleRequests() // Start handling HTTP requests
}
