package auth

import (
	"fmt"
	"net/http"

	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/api/responses"
	"github.com/labstack/echo/v4"
)

func handleRegisterUser (c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		req = &struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Name     string `json:"name"`
		}{}

	)

	fmt.Println(app)

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, responses.BadRequestResponse)
		return err
	}

	return nil
}

func handleSignIn (c echo.Context) error {
	return nil
}

func handleSignOut (c echo.Context) error {
	return nil
}