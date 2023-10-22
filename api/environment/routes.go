package environment

import "github.com/labstack/echo/v4"

func EnvironmentRoutes (g *echo.Group) {
	env := g.Group("/:orgUUID/envs")

	env.POST("/create", handleCreateEnvironments)
	env.GET("/", handleGetEnvironments)
	env.POST("/update/:envUUID", handleUpdateEnvironment)
	env.DELETE("/:envUUID", handleRemoveEnvironments)
}
