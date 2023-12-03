package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/api/responses"
	"github.com/chtushar/toggler/configs"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	cfg *configs.Config
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
  }
  return nil
}

func InitHTTPServer(app *app.App, ctx context.Context) *echo.Echo {
	srv := echo.New()
	srv.Validator = &CustomValidator{validator: validator.New()}
	// Passing the app instance to all the handlers
	// Helpful for using the database connection and logger
	srv.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("app", app)
			return next(c)
		}
	})

	cfg = configs.Get()
	// Initialize all the API handlers
	initHTTPHandler(srv)
	go func() {
		app.Log.Fatal(srv.Start(fmt.Sprintf(":%d", app.Port)))
	}()

	<- ctx.Done()
		fmt.Println("shutting down server")
		srv.Close()

	return srv
}