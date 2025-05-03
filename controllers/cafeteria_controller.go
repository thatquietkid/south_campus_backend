package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/config"
	"github.com/thatquietkid/south_campus_backend/models"
)

// GET all items
func GetCafeteriaItems(c echo.Context) error {
	var items []models.CafeteriaItem
	config.DB.Find(&items)
	return c.JSON(http.StatusOK, items)
}

// POST new item
func CreateCafeteriaItem(c echo.Context) error {
	var item models.CafeteriaItem
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid data"})
	}
	config.DB.Create(&item)
	return c.JSON(http.StatusCreated, item)
}

// DELETE item by ID
func DeleteCafeteriaItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}
	config.DB.Delete(&models.CafeteriaItem{}, id)
	return c.NoContent(http.StatusNoContent)
}
