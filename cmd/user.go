package cmd

import (
	"net/http"

	"github.com/chtushar/toggler/db/queries"
	"github.com/labstack/echo/v4"
)

func handleGetAllUsers(c echo.Context) error {
	var (
		app = c.Get("app").(*App)
	)

	type resType struct {
		Id            int32            `json:"id"`
		Name          string           `json:"name"`
		Email         string           `json:"email"`
		EmailVerified bool             `json:"email_verified"`
		Role          queries.UserRole `json:"role"`
	}

	users, err := app.q.GetAllUsers(c.Request().Context())

	if err != nil {
		app.log.Println("Failed to get all users", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	response := make([]resType, len(users))

	for i, user := range users {
		response[i] = resType{
			Id:            user.ID,
			Name:          user.Name,
			Email:         user.Email,
			EmailVerified: user.EmailVerified,
			Role:          user.Role,
		}
	}

	c.JSON(http.StatusOK, responseType{true, response, nil})
	return nil
}

func handleGetUserByID(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}

func handleGetUserByEmail(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}

func handleCreateUser(c echo.Context) error {
	var (
		app = c.Get("app").(*App)
		req = &struct {
			Name     string           `json:"name"`
			Email    string           `json:"email"`
			Password string           `json:"password"`
			Role     queries.UserRole `json:"role"`
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

	user, err := app.q.CreateUser(c.Request().Context(), queries.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
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

func handleUpdateUser(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}

func handleDeleteUser(c echo.Context) error {
	c.JSON(200, "ok")
	return nil
}
