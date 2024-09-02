package config

import (
	"database/sql" // Importing the sql package for database operations
	"fmt"          // Importing the fmt package for formatted I/O
	"log"          // Importing the log package for logging errors

	_ "github.com/go-sql-driver/mysql" // Importing the MySQL driver anonymously
)

var db *sql.DB // Declare a global variable to hold the database connection

// InitDB initializes the database connection
func InitDB() {
	var err error
	// Open a connection to the MySQL database using the provided DSN (Data Source Name)
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/employee_db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err) // Log and exit if connection fails
	}

	// Ping the database to verify the connection is established
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err) // Log and exit if ping fails
	}

	fmt.Println("Connected to database") // Confirm successful connection
}

// GetDB returns the database connection instance
func GetDB() *sql.DB {
	return db // Return the global database connection
}
