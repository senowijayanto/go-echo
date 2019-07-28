package models

// Employee struct
type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    int    `json:"age"`
}

// Employees is list employee
type Employees []Employee
