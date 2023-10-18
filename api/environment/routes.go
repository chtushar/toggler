package environment

import "github.com/labstack/echo/v4"

func EnvironmentRoutes (g *echo.Group) {
	env := g.Group("/envs")

	env.POST("/add", handleAddEnvironments)
}