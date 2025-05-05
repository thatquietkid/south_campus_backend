package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/config"
	"github.com/thatquietkid/south_campus_backend/models"
)

func GetAllHostels(c echo.Context) error {
	var hostels []models.Hostel

	if err := config.DB.Find(&hostels).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch hostel data",
		})
	}

	return c.JSON(http.StatusOK, hostels)
}