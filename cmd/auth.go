package cmd

import (
	"github.com/labstack/echo/v4"
)

func handleAddAdmin(c echo.Context) error {
	app := c.Get("app").(*App)

	count, err := app.q.CountUsers(c.Request().Context())

	if err != nil {
		app.log.Println("Failed to count users", err)
		return err
	}

	c.JSON(200, count)
	return nil
}

func handleLogin(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}

func handleLogout(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}
