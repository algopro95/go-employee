package models

import (
	"crud-app/config" // Importing the config package to access the database connection
	"database/sql"    // Importing the sql package for database operations
	"errors"          // Importing the errors package for error handling
)

// Employee struct represents the employee data model
type Employee struct {
	EmployeeID string // Unique identifier for the employee
	FullName   string // Full name of the employee
	Address    string // Address of the employee
}

// GetAllEmployees retrieves all employees from the database
func GetAllEmployees() ([]Employee, error) {
	db := config.GetDB() // Get the database connection

	// Execute the SQL query to select all employees
	rows, err := db.Query("SELECT employee_id, full_name, address FROM employees")
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	defer rows.Close() // Ensure that rows are closed after usage

	var employees []Employee

	// Iterate over the result set and scan each row into an Employee struct
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.EmployeeID, &emp.FullName, &emp.Address); err != nil {
			return nil, err // Return an error if scanning fails
		}
		employees = append(employees, emp) // Add the employee to the slice
	}

	return employees, nil // Return the list of employees and no error
}

// GetEmployeeByID retrieves an employee by their ID from the database
func GetEmployeeByID(id string) (Employee, error) {
	db := config.GetDB() // Get the database connection

	var emp Employee
	// Execute the SQL query to select an employee by ID
	err := db.QueryRow("SELECT employee_id, full_name, address FROM employees WHERE employee_id = ?", id).Scan(&emp.EmployeeID, &emp.FullName, &emp.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return emp, errors.New("employee not found") // Return a custom error if the employee is not found
		}
		return emp, err // Return an error if the query fails
	}

	return emp, nil // Return the found employee and no error
}

// AddEmployee adds a new employee to the database
func AddEmployee(emp Employee) error {
	// Validate that all fields are provided
	if emp.EmployeeID == "" || emp.FullName == "" || emp.Address == "" {
		return errors.New("all fields are required") // Return an error if any field is missing
	}

	// Check if the employee ID already exists
	existingEmp, err := GetEmployeeByID(emp.EmployeeID)
	if err == nil && existingEmp.EmployeeID != "" {
		return errors.New("employee ID already exists") // Return an error if the employee ID is a duplicate
	}

	db := config.GetDB() // Get the database connection

	// Execute the SQL query to insert a new employee
	_, err = db.Exec("INSERT INTO employees (employee_id, full_name, address) VALUES (?, ?, ?)", emp.EmployeeID, emp.FullName, emp.Address)
	if err != nil {
		return err // Return an error if the insertion fails
	}

	return nil // Return no error if the employee was added successfully
}

// UpdateEmployee updates an existing employee's data in the database
func UpdateEmployee(emp Employee) error {
	db := config.GetDB() // Get the database connection

	// Execute the SQL query to update the employee's data
	_, err := db.Exec("UPDATE employees SET full_name = ?, address = ? WHERE employee_id = ?", emp.FullName, emp.Address, emp.EmployeeID)
	if err != nil {
		return err // Return an error if the update fails
	}

	return nil // Return no error if the employee was updated successfully
}

// DeleteEmployee removes an employee from the database by their ID
func DeleteEmployee(id string) error {
	db := config.GetDB() // Get the database connection

	// Execute the SQL query to delete the employee by ID
	_, err := db.Exec("DELETE FROM employees WHERE employee_id = ?", id)
	if err != nil {
		return err // Return an error if the deletion fails
	}

	return nil // Return no error if the employee was deleted successfully
}
