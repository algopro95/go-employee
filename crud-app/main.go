package main

import (
	"crud-app/config"      // Importing the database configuration package
	"crud-app/controllers" // Importing the controllers package to handle business logic
	"fmt"                  // Importing the fmt package for formatted I/O
)

// The main function is the entry point of the Go application
func main() {
	// Initialize the database connection
	config.InitDB()

	// Add some sample employees to the database
	controllers.AddEmployee("1001", "Budi", "Jakarta")
	controllers.AddEmployee("1002", "Adi", "Jakarta")
	controllers.AddEmployee("1003", "Muhammad", "Tangerang")

	// Display all employees
	fmt.Println("All employees:")
	controllers.ShowAllEmployees()

	// Update the address of employee with ID "1001"
	fmt.Println("Updating employee 1001:")
	controllers.UpdateEmployee("1001", "Budi", "Bandung")
	controllers.ShowAllEmployees()

	// Delete the employee with ID "1002"
	fmt.Println("Deleting employee 1002:")
	controllers.DeleteEmployee("1002")
	controllers.ShowAllEmployees()
}
