package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, UnauthorizedResponse)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, UnauthorizedResponse)
		}

		return next(c)
	}
}

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

	hash, err := utils.HashPassword(req.Password)

	if err != nil {
		app.log.Println("Failed to hash password", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	user, err := app.q.CreateUser(c.Request().Context(), queries.CreateUserParams{
		Name:          req.Name,
		Email:         req.Email,
		EmailVerified: true,
		Password:      hash,
		Role:          queries.UserRoleAdmin,
	})

	if err != nil {
		app.log.Println("Failed to create user", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	token, err := generateToken(user.ID, user.Email, user.Name, user.Role)

	if err != nil {
		app.log.Println("Failed to generate token", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	fmt.Println(token)

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
