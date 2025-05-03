package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/config"
	"github.com/thatquietkid/south_campus_backend/models"
)

// GET /courses
func GetCourses(c echo.Context) error {
	var courses []models.Course
	config.DB.Find(&courses)
	return c.JSON(http.StatusOK, courses)
}

// GET /courses/:id
func GetCourseByID(c echo.Context) error {
	id := c.Param("id")
	var course models.Course
	result := config.DB.First(&course, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Course not found"})
	}
	return c.JSON(http.StatusOK, course)
}

// POST /courses
func CreateCourse(c echo.Context) error {
	var course models.Course
	if err := c.Bind(&course); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(&course).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create course"})
	}
	return c.JSON(http.StatusCreated, course)
}

// DELETE /courses/:id
func DeleteCourse(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Course{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete course"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Course deleted successfully"})
}
