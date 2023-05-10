package cmd

import (
	"fmt"
	"net/http"

	"github.com/chtushar/toggler/dashboard"
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

	// Echo group for protected routes
	g := e.Group("/api")
	g.Use(authMiddleware)

	// ...
	e.GET("/api/healthcheck", handleHealthCheck)

	// Auth
	e.POST("/api/add_admin", handleAddAdmin)
	e.POST("/api/login", handleLogin)
	g.POST("/logout", handleLogout)

	// Users
	g.GET("/api/get_user/:id", handleGetUserByID)
	g.GET("/api/get_user_by_email/:email", handleGetUserByEmail)
	g.GET("/api/get_users", handleGetAllUsers)
	g.POST("/api/create_user", handleCreateUser)
	g.PUT("/api/update_user/:id", handleUpdateUser)
	g.DELETE("/api/delete_user/:id", handleDeleteUser)

	fmt.Println("Initialized HTTP handlers")
}
