package organization

import "github.com/labstack/echo/v4"

func OrganizationRoutes(g *echo.Group) *echo.Group {
	org := g.Group("/organization")

	org.POST("/create", handleCreateOrganization)

	org_access := org.Group("")
	org_access.Use(CheckOrgAccessMiddleware)

	org_access.POST("/:orgUUID/add_members", handleAddOrganizationMembers)
	org_access.GET("/:orgUUID", handleGetOrganization)

	return org
}
