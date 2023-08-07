package cmd

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func handleGetUser(c echo.Context) error {
	var (
		user = c.Get("user").(*jwt.Token)
	)

	c.JSON(http.StatusOK, responseType{true, user.Claims.(jwt.MapClaims), nil})
	return nil
}
