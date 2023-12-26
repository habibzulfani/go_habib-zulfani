package http

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

var (
	secretKey = []byte("mom-rouy")
)

// SetJWTMiddleware sets up JWT middleware
func SetJWTMiddleware(e *echo.Echo) {
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: secretKey,
	}))
}

// AdminMiddleware is a middleware to check for admin role in the JWT claims
func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		return next(c)
	}
}
