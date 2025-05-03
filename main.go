package main

import (
	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/config"
	"github.com/thatquietkid/south_campus_backend/routes"
)

func main() {
	e := echo.New()

	config.ConnectDatabase()

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
