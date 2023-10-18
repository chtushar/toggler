package cmd

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/chtushar/toggler/app"
	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
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

func handleAddUser(c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		req = &struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Name     string `json:"name"`
		}{}
	)

	// Response type
	type resType struct {
		Id            int32            	`json:"id"`
		Uuid		  uuid.NullUUID		`json:"uuid"`
		Name          string           	`json:"name"`
		Email         string  			`json:"email"`
		EmailVerified bool             	`json:"email_verified"`
	}

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestResponse)
		return err
	}

	tx, err := app.DbConn.Begin(c.Request().Context())

	if err != nil {
		app.Log.Println("Failed to add member", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	defer tx.Rollback(c.Request().Context())
	qtx := app.Q.WithTx(tx)

	exists, err := qtx.CheckIfUserExists(c.Request().Context(), &req.Email)
	
	if err != nil {
		app.Log.Println("Failed to create user", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	// Create the user
	var user queries.User;


	// Hash the password
	hash, err := utils.HashPassword(req.Password)

	if err != nil {
		app.Log.Println("Failed to hash password", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	if exists {
		if err != nil {
			app.Log.Println("Failed to create user", err)
			c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
			return err
		}
		
		if user.EmailVerified {
			app.Log.Println("User already exists", err)
			c.JSON(http.StatusUnauthorized, UnauthorizedResponse)
			return fmt.Errorf("the user already exists")
		}

		user, err = qtx.UpdateUser(c.Request().Context(), queries.UpdateUserParams{
			Name: &req.Name,
			Email: &req.Email,
			EmailVerified: true,
		})

		if err != nil {
			app.Log.Println("Failed to create user", err)
			c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
			return err
		}

		qtx.UpdateUserPassword(c.Request().Context(), queries.UpdateUserPasswordParams{
			ID: user.ID,
			Password: &hash,	
		})
	} else {
		user, err = qtx.CreateUser(c.Request().Context(), queries.CreateUserParams{
			Name:          &req.Name,
			Email:         &req.Email,
			EmailVerified: true,
			Password:      &hash,
		})
	}

	if err != nil {
		app.Log.Println("Failed to create user", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	tx.Commit(c.Request().Context())

	if err != nil {
		var pgErr *pgconn.PgError
		errors.As(err, &pgErr)
		
		app.Log.Println("Failed to create user", err)

		if pgErr.Code == pgerrcode.UniqueViolation {
			c.JSON(http.StatusNotAcceptable, responseType{false, nil, &errorWrap{
				"User already exists",
				http.StatusNotAcceptable,
				nil,
			}})
			return err	
		}

		app.Log.Println("Failed to create user", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	// Generate JWT token
	token, err := generateToken(user.ID, user.Uuid, *user.Email, *user.Name)

	if err != nil {
		app.Log.Println("Failed to generate token", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	// Write token to cookie
	writeAuthTokenToCookie(c, token)

	c.JSON(http.StatusOK, responseType{true, user, nil})
	return nil
}

func handleLogin(c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		req = &struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
	)

	type resType struct {
		Id            int32            `json:"id"`
		Uuid		  string   			`json:"uuid"`
		Name          string           `json:"name"`
		Email         string   		   `json:"email"`
		EmailVerified bool             `json:"email_verified"`
	}

	if err := c.Bind(req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, BadRequestResponse)
		return err
	}

	user, err := app.Q.GetUserByEmail(c.Request().Context(), &req.Email)

	if err != nil {
		app.Log.Println("Failed to get user", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	if !utils.CheckPasswordHash(req.Password, *user.Password) {
		c.JSON(http.StatusUnauthorized, UnauthorizedResponse)
		return nil
	}

	token, err := generateToken(user.ID, user.Uuid, *user.Email, *user.Name)

	if err != nil {
		app.Log.Println("Failed to generate token", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	writeAuthTokenToCookie(c, token)

	response := resType{
		Id:            user.ID,
		Uuid: 		   user.Uuid,
		Name:          *user.Name,
		Email:         *user.Email,
		EmailVerified: user.EmailVerified,
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
