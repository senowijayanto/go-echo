package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq" // Postgresql driver
)

func main() {
	var err error
	db, err := sql.Open("postgres", "user=root password=secret dbname=books_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	// Echo instance
	e := echo.New()

	// GET all employee
	e.GET("/employee", func(c echo.Context) error {
		query := "SELECT id, name, salary, age FROM employees ORDER BY id"

		employees := Employees{}

		rows, err := db.Query(query)

		if err != nil {
			return err
		}

		defer rows.Close()

		for rows.Next() {
			employee := Employee{}

			err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age)

			if err != nil {
				return err
			}
			employees = append(employees, employee)
		}
		return c.JSON(http.StatusOK, employees)
	})

	// GET employee by id
	e.GET("/employee/:id", func(c echo.Context) error {
		id := c.Param("id")

		query := "SELECT id, name, salary, age FROM employees WHERE id = $1 ORDER BY id"

		employee := Employee{}

		statement, err := db.Prepare(query)

		if err != nil {
			return err
		}

		defer statement.Close()

		err = statement.QueryRow(id).Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age)

		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, &employee)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
