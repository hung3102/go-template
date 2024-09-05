package main

import (
	"gcim/example/config"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Load .env file
	config.LoadEnv()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1313"))
}
