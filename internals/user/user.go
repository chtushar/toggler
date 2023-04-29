package user

import (
	"github.com/labstack/echo/v4"
)

func HandleGetUsers(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}

func HandleGetUserByID(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}

func HandleGetUserByEmail(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}

func HandleCreateUser(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}

func HandleUpdateUser(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}

func HandleDeleteUser(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}
