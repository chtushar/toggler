package cmd

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handleHealthCheck is a healthcheck endpoint that returns a 200 response.
func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, responseType{true, "ok", nil})
}

func initHTTPHandler(e *echo.Echo, app *App) {
	fmt.Println("Initializing HTTP handlers")
	// ...
	e.GET("/api/healthcheck", handleHealthCheck)

	// Auth
	e.POST("/api/add_admin", handleAddAdmin)
	e.POST("/api/login", handleLogin)
	e.POST("/api/logout", handleLogout)

	// Users
	e.GET("/api/get_user/:id", handleGetUserByID)
	e.GET("/api/get_user_by_email/:email", handleGetUserByEmail)
	e.GET("/api/get_users", handleGetAllUsers)
	e.POST("/api/create_user", handleCreateUser)
	e.PUT("/api/update_user/:id", handleUpdateUser)
	e.DELETE("/api/delete_user/:id", handleDeleteUser)

	fmt.Println("Initialized HTTP handlers")
}
