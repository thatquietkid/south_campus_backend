package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/config"
	"github.com/thatquietkid/south_campus_backend/models"
)

func GetAllBusRoutes(c echo.Context) error {
	var busRoutes []models.BusRoute

	if err := config.DB.Find(&busRoutes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch bus route data",
		})
	}

	return c.JSON(http.StatusOK, busRoutes)
}