package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// PostgreSQL connection details (Dosmailova Dinara)
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func OpenNewDBConnection(dsn string) (*sql.DB, error) {
	// Open a connection to the database (Dosmailova Dinara)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the connection is working (Dosmailova Dinara)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INT
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
		return err
	}
	fmt.Println("Table created successfully!")
	return nil
}

func insertUser(db *sql.DB, name string, age int) error {
	query := `
	INSERT INTO users (name, age)
	VALUES ($1, $2);`

	_, err := db.Exec(query, name, age)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}
	fmt.Printf("User %s inserted successfully!\n", name)
	return nil
}

func getUsers(db *sql.DB) {
	query := `
	SELECT id, name, age
	FROM users;`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error querying users: %v", err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error after iteration: %v", err)
	}
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// connecting to db (Dosmailova Dinara)
	db, err := OpenNewDBConnection(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// defer - task execution at the end of the function
	defer db.Close()

	err = createTable(db)
	if err != nil {
		return
	}

	err = insertUser(db, "Dinara", 22)
	if err != nil {
		return
	}

	getUsers(db)
}
