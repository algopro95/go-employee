package controllers

import (
	"crud-app/models" // Importing the models package to interact with the data layer
	"crud-app/views"  // Importing the views package to handle the presentation of data
	"fmt"             // Importing the fmt package for formatted I/O
)

// ShowAllEmployees retrieves all employees and displays them using the view
func ShowAllEmployees() {
	employees, err := models.GetAllEmployees() // Fetch all employees from the model
	if err != nil {
		fmt.Println("Error fetching employees:", err) // Handle and display any errors
		return
	}

	views.DisplayEmployeeList(employees) // Pass the employee data to the view for display
}

// ShowEmployeeByID retrieves a single employee by ID and displays their details using the view
func ShowEmployeeByID(id string) {
	employee, err := models.GetEmployeeByID(id) // Fetch the employee by ID from the model
	if err != nil {
		fmt.Println("Error fetching employee:", err) // Handle and display any errors
		return
	}

	views.DisplayEmployeeDetails(employee) // Pass the employee data to the view for display
}

// AddEmployee creates a new employee record in the database
func AddEmployee(id, name, address string) {
	// Create a new Employee struct with the provided data
	emp := models.Employee{
		EmployeeID: id,
		FullName:   name,
		Address:    address,
	}

	// Call the model's AddEmployee function to insert the new employee
	err := models.AddEmployee(emp)
	if err != nil {
		fmt.Println("Error adding employee:", err) // Handle and display any errors
	} else {
		fmt.Println("Employee added successfully") // Confirm successful addition
	}
}

// UpdateEmployee updates an existing employee's information in the database
func UpdateEmployee(id, name, address string) {
	// Create an Employee struct with the updated data
	emp := models.Employee{
		EmployeeID: id,
		FullName:   name,
		Address:    address,
	}

	// Call the model's UpdateEmployee function to update the employee's information
	err := models.UpdateEmployee(emp)
	if err != nil {
		fmt.Println("Error updating employee:", err) // Handle and display any errors
	} else {
		fmt.Println("Employee updated successfully") // Confirm successful update
	}
}

// DeleteEmployee removes an employee record from the database by ID
func DeleteEmployee(id string) {
	// Call the model's DeleteEmployee function to delete the employee by ID
	err := models.DeleteEmployee(id)
	if err != nil {
		fmt.Println("Error deleting employee:", err) // Handle and display any errors
	} else {
		fmt.Println("Employee deleted successfully") // Confirm successful deletion
	}
}
