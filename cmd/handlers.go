package cmd

import (
	"fmt"
	"net/http"

	"github.com/chtushar/toggler/internals/user"
	"github.com/labstack/echo/v4"
)

type okResp struct {
	Data interface{} `json:"data"`
}

// handleHealthCheck is a healthcheck endpoint that returns a 200 response.
func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, okResp{true})
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
	e.GET("/api/get_user/:id", user.HandleGetUserByID)
	e.GET("/api/get_user_by_email/:email", user.HandleGetUserByEmail)
	e.GET("/api/get_users", user.HandleGetUsers)
	e.POST("/api/create_user", user.HandleCreateUser)
	e.PUT("/api/update_user/:id", user.HandleUpdateUser)
	e.DELETE("/api/delete_user/:id", user.HandleDeleteUser)

	fmt.Println("Initialized HTTP handlers")
}
