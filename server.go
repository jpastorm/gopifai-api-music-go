package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"wopifai/src/routes"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	routes.SetUpRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}