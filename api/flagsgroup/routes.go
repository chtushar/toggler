package flagsgroup

import "github.com/labstack/echo/v4"

func FlagGroupsRoutes (g *echo.Group) {
	flagsgroup := g.Group("/flagsgroup")

	flagsgroup.POST("/create", handleCreateFlagsGroup)
	flagsgroup.GET("/:fgUUID", handleGetFlagsGroup)
	flagsgroup.DELETE("/:fgUUID", handleDeleteFlagsGroup)
}