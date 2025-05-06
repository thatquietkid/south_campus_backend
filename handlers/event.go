package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/config"
	"github.com/thatquietkid/south_campus_backend/models"
)

// GetAllEvents fetches all events from the database
func GetAllEvents(c echo.Context) error {
	var events []models.Event

	if err := config.DB.Find(&events).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch event data",
		})
	}

	return c.JSON(http.StatusOK, events)
}