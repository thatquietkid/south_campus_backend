package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/handlers"
	"github.com/thatquietkid/south_campus_backend/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	// Public
	e.POST("/login", handlers.Login)
	e.POST("/register", handlers.Register)

	// Group that requires a valid JWT
	r := e.Group("")
	r.Use(middleware.JWTAuth)

	// Admin‐only subgroup
	admin := r.Group("")
	admin.Use(middleware.AdminOnly)

	// Public GET routes
	e.GET("/", handlers.Welcome)
	
}
