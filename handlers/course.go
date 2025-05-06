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

// GET /course-attendance
func GetCourseAttendance(c echo.Context) error {
	var courseAttendance []models.CourseAttendance
	err := config.DB.Raw(`
		SELECT c.title, c.code, a.attendance_percentage
		FROM courses c
		JOIN attendance a ON c.id = a.course_id
	`).Scan(&courseAttendance).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch course attendance"})
	}

	return c.JSON(http.StatusOK, courseAttendance)
}

// GET /course-syllabus/:code
func GetCourseSyllabusByCode(c echo.Context) error {
	code := c.Param("code")
	var syllabus []models.CourseSyllabus

	err := config.DB.Raw(`
		SELECT c.title, s.unit_title, s.topic
		FROM courses c
		JOIN syllabus s ON c.id = s.course_id
		WHERE c.code = ?
	`, code).Scan(&syllabus).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch syllabus"})
	}

	if len(syllabus) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "No syllabus found for the given course code"})
	}

	return c.JSON(http.StatusOK, syllabus)
}
