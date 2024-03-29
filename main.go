package main

import (
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq" // Postgresql driver
)

func main() {
	// Echo instance
	e := echo.New()

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
