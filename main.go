package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Employee struct {
	ID            int64
	UserName      string
	FirstName     string
	LastName      string
	Sex           string
	Email         string
	Designation   string
	ContactNumber int64
	BloodGroup    string
	TeamName      string
	DateOfJoining string
	Home          string
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "test",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	// Get Employee by username
	employee, err := employeeByUsername("mgs820")
	if err != nil {
		log.Fatal(err)
	}
	// Get Employee by teamname
	employee, err := employeeByTeamName("Destination")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Employee found: %v\n", employee)
	// Get Employee by id
	emp, err := employeeByID(6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Employee found: %v\n", emp)

	// Insert New Employee
	empId, err := addEmployee(Employee{
		UserName:      "sdy201",
		FirstName:     "Sandhya",
		LastName:      "Yadav",
		Sex:           "Female",
		Email:         "sandhyayadav@gmail.com",
		Designation:   "Drag and Drop",
		ContactNumber: 7005457545,
		BloodGroup:    "O+",
		TeamName:      "Customer Resolution",
		DateOfJoining: "2020-10-08",
		Home:          "Dawrka Delhi",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added employee: %v\n", empId)
}

func employeeByUsername(username string) ([]Employee, error) {

	var employee []Employee

	rows, err := db.Query("SELECT * FROM employee WHERE username = ?", username)
	if err != nil {
		return nil, fmt.Errorf("employeeByUsername %q: %v", username, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.UserName, &emp.FirstName, &emp.LastName, &emp.Sex, &emp.Email, &emp.Designation, &emp.ContactNumber, &emp.DateOfJoining, &emp.BloodGroup, &emp.TeamName, &emp.Home); err != nil {
			return nil, fmt.Errorf("employeeByUsername %q: %v", username, err)
		}
		employee = append(employee, emp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("employeeByUsername %q: %v", username, err)
	}
	return employee, nil
}

func employeeByID(id int64) (Employee, error) {
	var emp Employee

	row := db.QueryRow("SELECT * FROM employee WHERE id = ?", id)
	if err := row.Scan(&emp.ID, &emp.UserName, &emp.FirstName, &emp.LastName, &emp.Sex, &emp.Email, &emp.Designation, &emp.ContactNumber, &emp.DateOfJoining, &emp.BloodGroup, &emp.TeamName, &emp.Home); err != nil {
		if err == sql.ErrNoRows {
			return emp, fmt.Errorf("employeeByID %d: no such employee", id)
		}
		return emp, fmt.Errorf("employeeByID %d: %v", id, err)
	}
	return emp, nil
}
func employeeByTeamName(teamName string) ([]Employee, error) {

	var employee []Employee

	rows, err := db.Query("SELECT * FROM employee WHERE teamName = ?", teamName)
	if err != nil {
		return nil, fmt.Errorf("employeeByTeamName %q: %v", teamName, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.UserName, &emp.FirstName, &emp.LastName, &emp.Sex, &emp.Email, &emp.Designation, &emp.ContactNumber, &emp.DateOfJoining, &emp.BloodGroup, &emp.TeamName, &emp.Home); err != nil {
			return nil, fmt.Errorf("employeeByUsername %q: %v", teamName, err)
		}
		employee = append(employee, emp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("employeeByUsername %q: %v", teamName, err)
	}
	return employee, nil
}

func addEmployee(emp Employee) (int64, error) {
	result, err := db.Exec(`INSERT INTO employee (username,firstName,
        lastName,
        sex,
        email,
        designation,
        contactNumber,
        bloodGroup,
        teamName,
        dateOfJoining,
        home) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, emp.UserName, emp.FirstName, emp.LastName, emp.Sex, emp.Email, emp.Designation, emp.ContactNumber, emp.BloodGroup, emp.TeamName, emp.DateOfJoining, emp.Home)
	if err != nil {
		return 0, fmt.Errorf("addEmployee: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addEmployee: %v", err)
	}
	return id, nil
}
func employeeByEmail(Email string) ([]Employee, error) {

	var employee []Employee

	rows, err := db.Query("SELECT * FROM employee WHERE Email = ?", Email)
	if err != nil {
		return nil, fmt.Errorf("employeeByEmail %q: %v", Email, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.Email, &emp.FirstName, &emp.LastName, &emp.Sex, &emp.Email, &emp.Designation, &emp.ContactNumber, &emp.DateOfJoining, &emp.BloodGroup, &emp.TeamName, &emp.Home); err != nil {
			return nil, fmt.Errorf("employeeByEmail %q: %v", Email, err)
		}
		employee = append(employee, emp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("employeeByEmail %q: %v", Email, err)
	}
	return employee, nil
}

func employeeByDateOfJoining(DateOfJoining string) ([]Employee, error) {

	var employee []Employee

	rows, err := db.Query("SELECT * FROM employee WHERE DateOfJoining = ?", DateOfJoining)
	if err != nil {
		return nil, fmt.Errorf("employeeByEmail %q: %v", DateOfJoining, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.Email, &emp.FirstName, &emp.LastName, &emp.Sex, &emp.Email, &emp.Designation, &emp.ContactNumber, &emp.DateOfJoining, &emp.BloodGroup, &emp.TeamName, &emp.Home); err != nil {
			return nil, fmt.Errorf("employeeByEmail %q: %v", Email, err)
		}
		employee = append(employee, emp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("employeeByEmail %q: %v", DateOfJoining, err)
	}
	return employee, nil
}
