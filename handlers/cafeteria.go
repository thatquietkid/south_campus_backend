package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/config"
	"github.com/thatquietkid/south_campus_backend/models"
)

// GET /cafeteria-items
func GetCafeteriaItems(c echo.Context) error {
	var items []models.CafeteriaItem
	config.DB.Find(&items)
	return c.JSON(http.StatusOK, items)
}

// POST /cafeteria-items
func CreateCafeteriaItem(c echo.Context) error {
	var item models.CafeteriaItem
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(&item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create item"})
	}
	return c.JSON(http.StatusCreated, item)
}

// DELETE /cafeteria-items/:id
func DeleteCafeteriaItem(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.CafeteriaItem{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete item"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Item deleted successfully"})
}
