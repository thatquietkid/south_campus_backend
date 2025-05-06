package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/config"
	"github.com/thatquietkid/south_campus_backend/models"
)

func GetAllComplaints(c echo.Context) error {
	var complaints []models.Complaint
	if err := config.DB.Find(&complaints).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch complaints",
		})
	}
	return c.JSON(http.StatusOK, complaints)
}

func CreateComplaint(c echo.Context) error {
	var complaint models.Complaint
	if err := c.Bind(&complaint); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid complaint data",
		})
	}

	// Default status to "Pending"
	if complaint.Status == "" {
		complaint.Status = "Pending"
	}

	if err := config.DB.Create(&complaint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create complaint",
		})
	}
	return c.JSON(http.StatusCreated, complaint)
}
