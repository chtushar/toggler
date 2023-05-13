package cmd

import (
	"fmt"
	"net/http"

	"github.com/chtushar/toggler/dashboard"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// handleHealthCheck is a healthcheck endpoint that returns a 200 response.
func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, responseType{true, "ok", nil})
}

func initHTTPHandler(e *echo.Echo, app *App) {
	fmt.Println("Initializing HTTP handlers")

	// Serve static files
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: dashboard.BuildHTTPFS(),
		HTML5:      true,
	}))

	// ...
	e.GET("/api/healthcheck", handleHealthCheck)

	// Auth
	e.POST("/api/add_admin", handleAddAdmin)
	e.POST("/api/login", handleLogin)

	// Echo group for protected routes
	g := e.Group("/api")
	g.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(cfg.JWTSecret),
		TokenLookup: "cookie:auth_token",
		ErrorHandler: func(c echo.Context, err error) error {
			app.log.Println("JWT Error:", err)
			return c.JSON(http.StatusUnauthorized, UnauthorizedResponse)
		},
	}))

	// Auth
	g.POST("/logout", handleLogout)

	// Users
	g.GET("/get_user/:id", handleGetUserByID)
	g.GET("/get_user_by_email/:email", handleGetUserByEmail)
	g.GET("/get_users", handleGetAllUsers)
	g.POST("/create_user", handleCreateUser)
	g.PUT("/update_user/:id", handleUpdateUser)
	g.DELETE("/delete_user/:id", handleDeleteUser)

	fmt.Println("Initialized HTTP handlers")
}
