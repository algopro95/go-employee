# Employee CRUD Application

This GoLang application provides a simple CRUD (Create, Read, Update, Delete) interface for managing employee records using MySQL.

## Prerequisites

Before you start, ensure you have the following installed:

- [GoLang](https://golang.org/dl/) (version 1.16 or later)
- [MySQL](https://dev.mysql.com/downloads/mysql/) (version 5.7 or later)
- A MySQL client or database management tool (e.g., phpMyAdmin)

## Setup Guide

### 1. Install GoLang

1. **Download GoLang:**
   - Visit the [GoLang Downloads page](https://golang.org/dl/) and download the installer for your operating system.

2. **Install GoLang:**
   - Follow the installation instructions provided for your OS.

3. **Verify Installation:**
   - Open a terminal and run:
     ```bash
     go version
     ```
   - Ensure the output displays the Go version installed.

### 2. Install MySQL

1. **Download MySQL:**
   - Visit the [MySQL Downloads page](https://dev.mysql.com/downloads/mysql/) and download the installer for your operating system.

2. **Install MySQL:**
   - Follow the installation instructions provided.

3. **Verify Installation:**
   - Open a terminal and run:
     ```bash
     mysql --version
     ```
   - Ensure the output displays the MySQL version installed.

### 3. Set Up the Database

1. **Create a Database:**
   - Open a MySQL client or use a tool like phpMyAdmin.
   - Run the following SQL command to create a new database:
     ```sql
     CREATE DATABASE employee_db;
     ```

2. **Import the Database Schema:**
   - Place the `employee_db.sql` file in your project directory.
   - Open a terminal and navigate to the directory where `employee_db.sql` is located.
   - Run the following command to import the SQL file:
     ```bash
     mysql -u root -p employee_db < employee_db.sql
     ```
   - Enter your MySQL root password when prompted.

### 4. Configure the Application

1. **Update Database Configuration:**
   - Open `crud-app/config/db.go`.
   - Modify the connection string in the `sql.Open` function if necessary:
     ```go
     db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/employee_db")
     ```

### 5. Build and Run the Application

1. **Navigate to the Project Directory:**
   - Open a terminal and navigate to the root of your Go project:
     ```bash
     cd path/to/crud-app
     ```

2. **Build the Application:**
   - Run the following command to build the Go application:
     ```bash
     go build -o crud-app
     ```

3. **Run the Application:**
   - Execute the application with:
     ```bash
     ./crud-app
     ```

### 6. Application Usage

The application performs the following actions:

- **Add Employee:**
  - Adds a new employee record to the database.

- **Show All Employees:**
  - Lists all employee records.

- **Update Employee:**
  - Updates the details of an existing employee.

- **Delete Employee:**
  - Removes an employee record from the database.

## Troubleshooting

- **Database Connection Issues:**
  - Ensure MySQL is running and accessible.
  - Verify that the database configuration in `crud-app/config/db.go` is correct.

- **Build Errors:**
  - Ensure all dependencies are correctly imported in your Go modules.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [GoLang](https://golang.org/) for the programming language.
- [MySQL](https://www.mysql.com/) for the database.
