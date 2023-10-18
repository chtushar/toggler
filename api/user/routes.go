package user

import "github.com/labstack/echo/v4"

func UserRoutes(g *echo.Group) {
	user := g.Group("/user")

	user.GET("/me", handleGetUser)
	user.POST("/:userId/update", handleUpdateUser)
}