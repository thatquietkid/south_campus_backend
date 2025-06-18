package handlers

import (
	"net/http"
	"time"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/thatquietkid/south_campus_backend/models"
	"github.com/thatquietkid/south_campus_backend/config"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

// Struct for holding login credentials
type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

var jwtSecret = []byte("your_secret_key")

func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	// Use jwt.MapClaims to create token
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func Register(c echo.Context) error {
    var user models.User
    if err := c.Bind(&user); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Could not hash password")
    }
    user.Password = string(hashedPassword)

    // Save user to database
    if err := config.DB.Create(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Could not register user")
    }

    response := map[string]string{
        "message": "User registered successfully",
    }
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Could not encode response")
    }

    return c.Blob(http.StatusCreated, "application/json", jsonResponse)
}