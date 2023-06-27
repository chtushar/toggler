package cmd

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func writeAuthTokenToCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	c.SetCookie(cookie)
}

func handleHasAdmin(c echo.Context) error {
	var (
		app = c.Get("app").(*App)
	)

	count, err := app.q.CountUsers(c.Request().Context())

	if err != nil {
		app.log.Println("Failed to count users", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	c.JSON(http.StatusOK, responseType{true, count > 0, nil})
	return nil
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

	// Response type
	type resType struct {
		Id            int32            `json:"id"`
		Uuid		  uuid.NullUUID	`json:"uuid"`
		Name          string           `json:"name"`
		Email         string  `json:"email"`
		EmailVerified bool             `json:"email_verified"`
		Role          queries.UserRole `json:"role"`
	}

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestResponse)
		return err
	}

	// Hash the password
	hash, err := utils.HashPassword(req.Password)

	if err != nil {
		app.log.Println("Failed to hash password", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	// Create the user with admin role
	user, err := app.q.CreateUser(c.Request().Context(), queries.CreateUserParams{
		Name:          req.Name,
		Email:         sql.NullString{ String: req.Email, Valid: true},
		EmailVerified: true,
		Password:      hash,
		Role:          queries.UserRoleAdmin,
	})

	if err != nil {
		app.log.Println("Failed to create user", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	// Generate JWT token
	token, err := generateToken(user.ID, user.Uuid, user.Email, user.Name, user.Role)

	if err != nil {
		app.log.Println("Failed to generate token", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	// Write token to cookie
	writeAuthTokenToCookie(c, token)

	response := resType{
		Id:            user.ID,
		Name:          user.Name,
		Email:         user.Email.String,
		EmailVerified: user.EmailVerified,
		Role:          user.Role,
	}

	c.JSON(http.StatusOK, responseType{true, response, nil})
	return nil
}

func handleLogin(c echo.Context) error {
	var (
		app = c.Get("app").(*App)
		req = &struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
	)

	type resType struct {
		Id            int32            `json:"id"`
		Uuid		  uuid.NullUUID	   `json:"uuid"`
		Name          string           `json:"name"`
		Email         string   		   `json:"email"`
		EmailVerified bool             `json:"email_verified"`
		Role          queries.UserRole `json:"role"`
	}

	if err := c.Bind(req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, BadRequestResponse)
		return err
	}

	user, err := app.q.GetUserByEmail(c.Request().Context(), sql.NullString{String: req.Email, Valid: true })

	if err != nil {
		app.log.Println("Failed to get user", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, UnauthorizedResponse)
		return nil
	}

	token, err := generateToken(user.ID, user.Uuid, user.Email, user.Name, user.Role)

	if err != nil {
		app.log.Println("Failed to generate token", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	writeAuthTokenToCookie(c, token)

	response := resType{
		Id:            user.ID,
		Uuid: 		   user.Uuid,
		Name:          user.Name,
		Email:         user.Email.String,
		EmailVerified: user.EmailVerified,
		Role:          user.Role,
	}

	c.JSON(http.StatusOK, responseType{true, response, nil})
	return nil
}

func handleLogout(c echo.Context) error {

	// Clear the auth token
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	c.SetCookie(cookie)

	c.JSON(200, "Logged Out")
	return nil
}
