package cmd

import (
	"net/http"
	"strconv"

	"github.com/chtushar/toggler/dashboard"
	"github.com/chtushar/toggler/db/queries"
	"github.com/golang-jwt/jwt/v4"
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

	v1_org_access := v1_protected.Group("")
	v1_org_access.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var  (
				app  = c.Get("app").(*App)
				user = c.Get("user").(*jwt.Token)
		)
			claims := user.Claims.(jwt.MapClaims)
			orgIdParam := c.Param("orgId")
			orgId, err := strconv.Atoi(orgIdParam)

			if err != nil {
				c.JSON(http.StatusBadRequest, BadRequestResponse)
				app.log.Println("Coulnd't get the org id")
				return err
			}			
			userId := int64(claims["id"].(float64))

			ok, err := app.q.DoesUserBelongToOrg(c.Request().Context(), queries.DoesUserBelongToOrgParams{
				OrgID: int64(orgId),
				UserID: userId,
			})

			if err != nil {
				c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
				app.log.Println("Failed to check the access of the org", err)
				return err
			}
			
			if !ok {
				c.JSON(http.StatusUnauthorized, UnauthorizedResponse)
			}
			
			c.Set("orgId", orgId)
			
			return next(c)
		}
	})

	// Organizations
	v1_protected.POST("/create_organization", handleCreateOrganization)
	v1_protected.GET("/get_user_organizations", handleGetUserOrganizations)
	v1_protected.POST("/update_organization", handleUpdateOrganization)
	v1_org_access.GET("/get_team_members/:orgId", handleGetOrganizationMembers)
	v1_org_access.POST("/add_team_member/:orgId", handleAddTeamMember)

	// Auth
	v1_protected.POST("/logout", handleLogout)

	// Users
	v1_protected.GET("/get_user", handleGetUser)

	// Projects
	v1_org_access.POST("/create_project/:orgId", handleCreateProject)
	v1_org_access.GET("/get_org_projects/:orgId", handleGetOrgProjects)

	// Environments
	v1_protected.GET("/get_project_environments/:projectId", handleGetProjectEnvironments)

	// Feature Flags
	v1_protected.POST("/create_feature_flag", handleCreateFeatureFlag)
	v1_protected.GET("/get_project_feature_flags/:projectId/:environmentId", handleGetProjectFeatureFlags)
	v1.GET("/get_flags", handleGetFlags)
}
