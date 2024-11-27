package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "exercise-two/docs"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	sqlDB    *sql.DB
	gormDB   *gorm.DB
	validate *validator.Validate
)

// Dinara Dosmailova
// User model for GORM
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null" validate:"required"`
	Age  int    `json:"age" gorm:"not null" validate:"required,gt=0"`
}

// @title Go REST API
// @version 1.0
// @description This API demonstrates CRUD operations using database/sql and GORM with Swagger documentation.
// @host localhost:8000
// @BasePath /
func main() {
	initSQLDB()
	initGORMDB()

	// Dinara Dosmailova
	// Initialize validator
	validate = validator.New()

	// Setup router
	r := mux.NewRouter()

	// Dinara Dosmailova
	// Define API routes
	r.HandleFunc("/users/sql", getUsersSQL).Methods("GET")
	r.HandleFunc("/user/sql", createUserSQL).Methods("POST")
	r.HandleFunc("/user/sql/{id}", updateUserSQL).Methods("PUT")
	r.HandleFunc("/user/sql/{id}", deleteUserSQL).Methods("DELETE")

	r.HandleFunc("/users/gorm", getUsersGORM).Methods("GET")
	r.HandleFunc("/user/gorm", createUserGORM).Methods("POST")
	r.HandleFunc("/user/gorm/{id}", updateUserGORM).Methods("PUT")
	r.HandleFunc("/user/gorm/{id}", deleteUserGORM).Methods("DELETE")

	// Dinara Dosmailova
	// Swagger documentation endpoint
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Dinara Dosmailova
	// Start the server
	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Dinara Dosmailova
// Initialize database connection using database/sql
func initSQLDB() {
	var err error
	psqlInfo := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	sqlDB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to database/sql: %v", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("Unable to ping database/sql: %v", err)
	}
	fmt.Println("Connected to database using database/sql")
}

// Dinara Dosmailova
// Initialize database connection using GORM
func initGORMDB() {
	var err error
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database using GORM: %v", err)
	}
	// Dinara Dosmailova
	gormDB.AutoMigrate(&User{})
	fmt.Println("Connected to database using GORM and AutoMigrate successful")
}

// @Summary Retrieve all users
// @Description Fetches a list of all users from the database using the database/sql package
// @Tags Users
// @Produce json
// @Success 200 {array} User
// @Failure 500 {string} string "Unable to query users"
// @Router /users/sql [get]
func getUsersSQL(w http.ResponseWriter, r *http.Request) {
	rows, err := sqlDB.Query("SELECT id, name, age FROM users")
	if err != nil {
		http.Error(w, "Unable to query users", http.StatusInternalServerError)
		return
	}
	// Dinara Dosmailova
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			http.Error(w, "Unable to scan user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	// Dinara Dosmailova
	json.NewEncoder(w).Encode(users)
}

// @Summary Create a new user
// @Description Adds a new user to the database using the database/sql package
// @Tags Users
// @Accept json
// @Produce json
// @Param user body User true "User to create"
// @Success 200 {object} User
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Error inserting user"
// @Router /user/sql [post]
func createUserSQL(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	} // Dinara Dosmailova
	if user.Name == "" || user.Age <= 0 {
		http.Error(w, "Name and age are required", http.StatusBadRequest)
		return
	}
	query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id"
	err := sqlDB.QueryRow(query, user.Name, user.Age).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	} // Dinara Dosmailova
	json.NewEncoder(w).Encode(user)
}

// @Summary Update a user
// @Description Updates an existing user's information in the database using the database/sql package
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body User true "Updated user data"
// @Success 200 {string} string "User updated successfully"
// @Failure 400 {string} string "Invalid user ID or request payload"
// @Failure 404 {string} string "User not found or no change in data"
// @Failure 500 {string} string "Error updating user"
// @Router /user/sql/{id} [put]
func updateUserSQL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	} // Dinara Dosmailova
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	query := "UPDATE users SET name = $1, age = $2 WHERE id = $3"
	result, err := sqlDB.Exec(query, user.Name, user.Age, id)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	} // Dinara Dosmailova
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "User not found or no change in data", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete a user
// @Description Deletes a user from the database using the database/sql package
// @Tags Users
// @Param id path int true "User ID"
// @Success 204 {string} string "No content"
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Error deleting user"
// @Router /user/sql/{id} [delete]
func deleteUserSQL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	query := "DELETE FROM users WHERE id = $1"
	// Dinara Dosmailova
	// Implementation of query responce
	result, err := sqlDB.Exec(query, id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}
	// Dinara Dosmailova
	//  Checks whether the user is deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// @Summary Retrieve all users (GORM)
// @Description Fetches a list of all users from the database using GORM
// @Tags Users
// @Produce json
// @Success 200 {array} User
// @Router /users/gorm [get]
func getUsersGORM(w http.ResponseWriter, r *http.Request) {
	var users []User
	gormDB.Find(&users)
	json.NewEncoder(w).Encode(users)
	// Dinara Dosmailova
}

// @Summary Create a new user (GORM)
// @Description Adds a new user to the database using GORM
// @Tags Users
// @Accept json
// @Produce json
// @Param user body User true "User to create"
// @Success 200 {object} User
// @Failure 400 {string} string "Invalid request payload or validation error"
// @Failure 500 {string} string "Error inserting user"
// @Router /user/gorm [post]
func createUserGORM(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Dinara Dosmailova
	if err := validate.Struct(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Dinara Dosmailova
	result := gormDB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// Dinara Dosmailova
	json.NewEncoder(w).Encode(user)
}

// @Summary Update a user (GORM)
// @Description Updates an existing user's information in the database using GORM
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body User true "Updated user data"
// @Success 200 {object} User
// @Failure 400 {string} string "Invalid request payload or validation error"
// @Failure 404 {string} string "User not found"
// @Router /user/gorm/{id} [put]
func updateUserGORM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user User
	// Dinara Dosmailova
	if gormDB.First(&user, id).Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Dinara Dosmailova
	if err := validate.Struct(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gormDB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// @Summary Delete a user (GORM)
// @Description Deletes a user from the database using GORM
// @Tags Users
// @Param id path int true "User ID"
// @Success 204 {string} string "No content"
// @Failure 404 {string} string "User not found"
// @Router /user/gorm/{id} [delete]
func deleteUserGORM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user User
	// Dinara Dosmailova
	if gormDB.First(&user, id).Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	gormDB.Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}
