package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/config"
	"github.com/thatquietkid/south_campus_backend/models"
)

func GetAllBusRoutes(c echo.Context) error {
	var routes []models.Route

	// Fetch basic route data
	if err := config.DB.Find(&routes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch bus routes",
		})
	}

	// Fetch and attach schedules for each route
	for i := range routes {
		var stops []models.Stop
		if err := config.DB.
			Table("stops").
			Where("route_number = ?", routes[i].Number).
			Find(&stops).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Failed to fetch schedule for route " + routes[i].Number,
			})
		}
		routes[i].Schedule = stops
	}

	return c.JSON(http.StatusOK, routes)
}
