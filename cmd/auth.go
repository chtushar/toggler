package cmd

import (
	"net/http"

	"github.com/chtushar/toggler/db/queries"
	"github.com/labstack/echo/v4"
)

func handleAddAdmin(c echo.Context) error {
	var (
		app = c.Get("app").(*App)
		req = &struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Name     string `json:"name"`
		}{}
	)

	type resType struct {
		Id            int32            `json:"id"`
		Name          string           `json:"name"`
		Email         string           `json:"email"`
		EmailVerified bool             `json:"email_verified"`
		Role          queries.UserRole `json:"role"`
	}

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestResponse)
		return err
	}

	count, err := app.q.CountUsers(c.Request().Context())

	if err != nil {
		app.log.Println("Failed to count users", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	if count > 0 {
		c.JSON(http.StatusForbidden, ForbiddenResponse)
		return nil
	}

	user, err := app.q.CreateUser(c.Request().Context(), queries.CreateUserParams{
		Name:          req.Name,
		Email:         req.Email,
		EmailVerified: true,
		Password:      req.Password,
		Role:          queries.UserRoleAdmin,
	})

	if err != nil {
		app.log.Println("Failed to create user", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	response := resType{
		Id:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		Role:          user.Role,
	}

	c.JSON(http.StatusOK, responseType{true, response, nil})
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
