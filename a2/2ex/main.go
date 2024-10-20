package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	Age  int
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Error in AutoMigrate: %v", err)
	}
	fmt.Println("Table auto-migrated successfully!")
}

func insertUser(db *gorm.DB, name string, age int) {
	user := User{Name: name, Age: age}
	result := db.Create(&user)

	if result.Error != nil {
		log.Fatalf("Error inserting user: %v", result.Error)
	}
	fmt.Printf("User %s inserted successfully with ID: %d!\n", user.Name, user.ID)
}

func getUsers(db *gorm.DB) {
	var users []User
	result := db.Find(&users)

	if result.Error != nil {
		log.Fatalf("Error querying users: %v", result.Error)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}

func main() {
	// PostgreSQL connection details
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	// Initialize the PostgreSQL connection with GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Successfully connected to PostgreSQL using GORM!")

	// Auto migrate the User struct to create/update the `users` table
	autoMigrate(db)

	// Insert users
	insertUser(db, "Alice", 30)
	insertUser(db, "Bob", 25)

	// Query and print all users
	getUsers(db)
}
