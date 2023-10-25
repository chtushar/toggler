package apikeys

import "github.com/labstack/echo/v4"

func APIKeysRoutes(g *echo.Group) {
	apiKeys := g.Group("/api_keys")

	apiKeys.GET("", handleGetOrgAPIKeys)
	apiKeys.POST("/create", handleCreateAPIKey)
	apiKeys.DELETE("/:apiKeyUUID", handleDeleteAPIKey)
}
