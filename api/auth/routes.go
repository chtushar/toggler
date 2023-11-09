package auth

import "github.com/labstack/echo/v4"

func AuthRoutes(g *echo.Group) {
	auth := g.Group("/auth")

	auth.POST("/register/", handleRegisterUser)
	auth.POST("/signin/", handleSignIn)
	auth.POST("/signout/", handleSignOut)

}
