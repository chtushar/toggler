package api

import (
	"net/http"

	"github.com/chtushar/toggler/api/apikeys"
	"github.com/chtushar/toggler/api/auth"
	"github.com/chtushar/toggler/api/flagsgroup"
	"github.com/chtushar/toggler/api/flagsgroupstate"
	"github.com/chtushar/toggler/api/folder"
	"github.com/chtushar/toggler/api/organization"
	"github.com/chtushar/toggler/api/responses"
	"github.com/chtushar/toggler/api/user"
	"github.com/chtushar/toggler/dashboard"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initHTTPHandler(e *echo.Echo) {
	e.Pre(middleware.AddTrailingSlash())
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

	// // Echo group for protected routes
	v1_protected := v1.Group("")
	v1_protected.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(cfg.JWTSecret),
		TokenLookup: "cookie:auth_token",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, responses.UnauthorizedResponse)
		},
	}))

	auth.AuthRoutes(v1)
	user.UserRoutes(v1_protected)
	org_access := organization.OrganizationRoutes(v1_protected)

	folder.FolderRoutes(org_access)
	apikeys.APIKeysRoutes(org_access)
	flagsgroup.FlagGroupsRoutes(org_access)
	flagsgroupstate.FlagsGroupStateRoutes(org_access)
}
