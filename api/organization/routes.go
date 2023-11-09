package organization

import "github.com/labstack/echo/v4"

func OrganizationRoutes(g *echo.Group) *echo.Group {
	org := g.Group("/organization")

	org.GET("/", handleGetUserOrganizations)
	org.POST("/create/", handleCreateOrganization)

	org_access := org.Group("")
	org_access.Use(CheckOrgAccessMiddleware)
	org_access.GET("/:orgUUID/", handleGetOrganization)
	org_access.POST("/:orgUUID/add_members/", handleAddOrganizationMembers)

	return org_access
}
