package cmd

import (
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
	
	// ...
	v1.GET("/healthcheck", handleHealthCheck)

	// Auth
	v1.POST("/add_user", handleAddUser)
	v1.POST("/login", handleLogin)

	// Echo group for protected routes
	v1_protected := v1.Group("")
	v1_protected.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(cfg.JWTSecret),
		TokenLookup: "cookie:auth_token",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, UnauthorizedResponse)
		},
	}))

	// Organizations
	v1_protected.POST("/create_organization", handleCreateOrganization)
	v1_protected.GET("/get_user_organizations", handleGetUserOrganizations)
	v1_protected.POST("/update_organization", handleUpdateOrganization)

	// Auth
	v1_protected.POST("/logout", handleLogout)

	// Users
	v1_protected.GET("/get_user", handleGetUser)
	v1_protected.GET("/get_user/:id", handleGetUserByID)

	// Projects
	v1_protected.POST("/create_project", handleCreateProject)
	v1_protected.GET("/get_org_projects/:orgId", handleGetOrgProjects)

	// Environments
	v1_protected.GET("/get_project_environments/:projectId", handleGetProjectEnvironments)

	// Feature Flags
	v1_protected.POST("/create_feature_flag", handleCreateFeatureFlag)
	v1_protected.GET("/get_project_feature_flags/:projectId/:environmentId", handleGetProjectFeatureFlags)
	v1.GET("/get_flags", handleGetFlags)
}
