package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thatquietkid/south_campus_backend/config"
	"github.com/thatquietkid/south_campus_backend/routes"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	config.ConnectDatabase()

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
