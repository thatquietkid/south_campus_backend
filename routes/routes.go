package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/handlers"
	"github.com/thatquietkid/south_campus_backend/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	// Public
	e.POST("/login", handlers.Login)

	// Group that requires a valid JWT
	r := e.Group("")
	r.Use(middleware.JWTAuth)

	// Admin‐only subgroup
	admin := r.Group("")
	admin.Use(middleware.AdminOnly)
	admin.POST("/courses", handlers.CreateCourse)
	admin.DELETE("/courses/:id", handlers.DeleteCourse)
	admin.POST("/cafeteria-items", handlers.CreateCafeteriaItem)
	admin.DELETE("/cafeteria-items/:id", handlers.DeleteCafeteriaItem)

	// Public GET routes
	e.GET("/", handlers.Welcome)
	e.GET("/courses", handlers.GetCourses)
	e.GET("/courses/:id", handlers.GetCourseByID)
	e.GET("/cafeteria-items", handlers.GetCafeteriaItems)
	e.GET("/hostels", handlers.GetAllHostels)
	e.GET("/busroutes", handlers.GetAllBusRoutes)
}
