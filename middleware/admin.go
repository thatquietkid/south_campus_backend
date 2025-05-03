package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := c.Get("user")
		if v == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing user context"})
		}

		tok := v.(*jwt.Token)
		claims := tok.Claims.(jwt.MapClaims)

		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			return c.JSON(http.StatusForbidden, echo.Map{"error": "Admin access required"})
		}

		return next(c)
	}
}
