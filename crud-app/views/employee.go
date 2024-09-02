package views

import (
	"crud-app/models"
	"fmt"
)

// DisplayEmployeeList prints the list of employees to the console.
func DisplayEmployeeList(employees []models.Employee) {
	fmt.Println("Employee List:")
	for _, emp := range employees {
		fmt.Printf("Employee ID: %s, Full Name: %s, Address: %s\n", emp.EmployeeID, emp.FullName, emp.Address)
	}
}

// DisplayEmployeeDetails prints details of a single employee.
func DisplayEmployeeDetails(emp models.Employee) {
	fmt.Printf("Employee Details:\nEmployee ID: %s\nFull Name: %s\nAddress: %s\n", emp.EmployeeID, emp.FullName, emp.Address)
}
