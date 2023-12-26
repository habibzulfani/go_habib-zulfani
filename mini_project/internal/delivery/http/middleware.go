package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetMiddleware sets the middleware for the Echo instance
func SetMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
}

