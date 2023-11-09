package flagsgroupstate

import "github.com/labstack/echo/v4"

func FlagsGroupStateRoutes(g *echo.Group) {
	fgs := g.Group("/:orgUUID/flagsgroup/:fgUUID/flagsgroupstate")

	fgs.GET("/", handleGetFlagsGroupState)
	fgs.POST("/update_json/", handleUpdateFlagsGroupStateJSON)
}
