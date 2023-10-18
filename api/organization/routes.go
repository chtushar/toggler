package organization

import "github.com/labstack/echo/v4"

func OrganizationRoutes(g *echo.Group) *echo.Group {
	org := g.Group("/organization")

	org.POST("/create", handleCreateOrganization)
	org.POST("/add_members", handleAddOrganizationMembers)
	org.GET("/:orgUUID", handleGetOrganization)

	return org
}
