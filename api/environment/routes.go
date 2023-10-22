package environment

import "github.com/labstack/echo/v4"

func EnvironmentRoutes (g *echo.Group) {
	env := g.Group("/:orgUUID/envs")

	env.GET("", handleGetEnvironments)
	env.POST("/create", handleCreateEnvironments)
	env.POST("/update/:envUUID", handleUpdateEnvironment)
}
