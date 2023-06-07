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

	// Check if the admin exists
	// If not, redirect to /register-admin

	// ...
	e.GET("/api/healthcheck", handleHealthCheck)
	e.GET("/api/has_admin", handleHasAdmin)

	// Auth
	e.POST("/api/add_admin", handleAddAdmin)
	e.POST("/api/login", handleLogin)

	// Echo group for protected routes
	g := e.Group("/api")
	g.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(cfg.JWTSecret),
		TokenLookup: "cookie:auth_token",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, UnauthorizedResponse)
		},
	}))

	// Auth
	g.POST("/logout", handleLogout)

	// Users
	g.GET("/get_user", handleGetUser)
	g.GET("/get_user/:id", handleGetUserByID)
	g.GET("/get_user_by_email/:email", handleGetUserByEmail)
	g.GET("/get_users", handleGetAllUsers)
	g.POST("/create_user", handleCreateUser)
	g.PUT("/update_user/:id", handleUpdateUser)
	g.DELETE("/delete_user/:id", handleDeleteUser)

	// Projects
	g.POST(("/create_project"), handleCreateProject)
	g.GET("/get_user_projects", handleGetUserProjects)

	fmt.Println("Initialized HTTP handlers")
}
