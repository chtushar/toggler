package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/chtushar/toggler/api/app"
	j "github.com/chtushar/toggler/api/jwt"
	"github.com/chtushar/toggler/api/responses"
	u "github.com/chtushar/toggler/api/user"
	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
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

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, responses.BadRequestResponse)
		return err
	}

	hash, err := utils.HashPassword(req.Password)

	if err != nil {
		app.Log.Println("Failed to hash password", err)
		c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		return err
	}


	tx, err := app.DbConn.Begin(c.Request().Context())

	if err != nil {
		app.Log.Println("Failed to add member", err)
		c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		return err
	}

	defer tx.Rollback(c.Request().Context())

	qtx := app.Q.WithTx(tx)

	user , err := qtx.CreateActiveUser(c.Request().Context(), queries.CreateActiveUserParams{
		Email: req.Email,
		Password: hash,
		Name: req.Name,
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				c.JSON(http.StatusBadRequest, responses.ErrorResponse(http.StatusBadRequest, "Email already exists"))
				return err
			}
		}
		c.JSON(http.StatusBadRequest, responses.InternalServerErrorResponse)
		return err
	}

	tx.Commit(c.Request().Context())

	token, err := j.GenerateToken(jwt.MapClaims{
		"uuid":  user.Uuid,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}, app.Jwt)

	if err != nil {
		app.Log.Println("Failed to generate token", err)
		c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		return err
	}

	writeAuthTokenToCookie(c, token)

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: u.UserNoPassword{
			Uuid: user.Uuid,
			Name: user.Name,
			Email: user.Email,
			EmailVerified: *user.EmailVerified,
			Active: *user.Active,
			CreatedAt: *user.CreatedAt,
		},
		Error: nil,
	})
	return nil
}

func handleSignIn (c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		req = &struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
	)

	if err := c.Bind(req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, responses.BadRequestResponse)
		return err
	}

	user, err := app.Q.GetUserByEmail(c.Request().Context(), req.Email)

	if err != nil {
		c.JSON(http.StatusForbidden, responses.ErrorResponse(http.StatusForbidden, "Incorrect Credentials"))
		return err
	}

	ok := utils.CheckPasswordHash(req.Password, user.Password)

	if !ok {
		c.JSON(http.StatusForbidden, responses.ErrorResponse(http.StatusForbidden, "Incorrect Credentials"))
		return err
	}

	token, err := j.GenerateToken(jwt.MapClaims{
		"uuid":  user.Uuid,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}, app.Jwt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		return err
	}

	writeAuthTokenToCookie(c, token)

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: u.UserNoPassword {
			Uuid: user.Uuid,
			Name: user.Name,
			Email: user.Email,
			EmailVerified: *user.EmailVerified,
			Active: *user.Active,
			CreatedAt: *user.CreatedAt,
		},
		Error: nil,
	})
	return nil
}

func handleSignOut (c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	c.SetCookie(cookie)

	c.JSON(200, responses.ResponseType{
		Success: true,
		Data: "Logged Out",
		Error: nil,
	})
	return nil
}