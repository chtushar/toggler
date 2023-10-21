package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/chtushar/toggler/api/app"
	j "github.com/chtushar/toggler/api/jwt"
	"github.com/chtushar/toggler/api/responses"
	u "github.com/chtushar/toggler/api/user"
	"github.com/chtushar/toggler/db"
	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/labstack/echo/v4"
)

func handleRegisterUser (c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		req = &struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required,min=5"`
			Name     string `json:"name" validate:"required,min=2"`
		}{}
	)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	hash, err := utils.HashPassword(req.Password)

	if err != nil {
		app.Log.Println("Failed to hash password", err)
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	user, err := db.WithDBTransaction[queries.User](app, c.Request().Context(), func(q *queries.Queries) (*queries.User, error) {
		user, err := q.CreateActiveUser(c.Request().Context(), queries.CreateActiveUserParams{
			Name: req.Name,
			Email: req.Email,
			Password: hash,
		})

		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == pgerrcode.UniqueViolation {
					return nil, echo.NewHTTPError(http.StatusBadRequest, responses.ErrorResponse(http.StatusBadRequest, "Email already exists"))
				}
			}
			return nil, echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
		}

		return &user, nil
	})

	if err != nil {
		return err
	}

	token, err := j.GenerateToken(jwt.MapClaims{
		"uuid":  user.Uuid,
		"email": user.Email,
		"name":  user.Name,
		"org": "",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}, app.Jwt)

	if err != nil {
		app.Log.Println("Failed to generate token", err)
		c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		return err
	}

	WriteAuthTokenToCookie(c, token)

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
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	user, err := app.Q.GetUserByEmail(c.Request().Context(), req.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, responses.ErrorResponse(http.StatusForbidden, "Incorrect Credentials"))
	}

	ok := utils.CheckPasswordHash(req.Password, user.Password)

	if !ok {
		return echo.NewHTTPError(http.StatusForbidden, responses.ErrorResponse(http.StatusForbidden, "Incorrect Credentials"))
	}

	orgs, err := app.Q.GetUserOrganizations(c.Request().Context(), user.Uuid)
	
	orgUUIDs := ""
	
	if err == nil {
		for i, o := range orgs {
			if i == 0 {
				orgUUIDs = o.Uuid
				continue
			}
			orgUUIDs = orgUUIDs + "," + o.Uuid

		}
	}
	
	token, err := j.GenerateToken(jwt.MapClaims{
		"uuid":  user.Uuid,
		"email": user.Email,
		"name":  user.Name,
		"orgs": 	 orgUUIDs,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}, app.Jwt)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	WriteAuthTokenToCookie(c, token)

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