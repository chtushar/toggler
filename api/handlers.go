package api

import (
	"net/http"

	"github.com/chtushar/toggler/dashboard"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initHTTPHandler(e *echo.Echo) {

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:9091"},
		AllowCredentials: true,
	}))
	
	// Middlewares
	// Serve static files
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: dashboard.BuildHTTPFS(),
		HTML5:      true,
	}))
	
	v1 := e.Group("/api/v1")
	
	// Echo group for protected routes
	v1_protected := v1.Group("")
	v1_protected.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(cfg.JWTSecret),
		TokenLookup: "cookie:auth_token",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, UnauthorizedResponse)
		},
	}))

}