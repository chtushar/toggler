package api

import (
	"fmt"
	"log"

	"github.com/chtushar/toggler/configs"
	"github.com/chtushar/toggler/db/queries"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

type App struct {
	Port   int
	Jwt    string
	DbConn *pgxpool.Pool
	Q      *queries.Queries
	Log    *log.Logger
}

var (
	cfg *configs.Config
)


func InitHTTPServer(app *App) *echo.Echo {
	srv := echo.New()

	// Passing the app instance to all the handlers
	// Helpful for using the database connection and logger
	srv.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("app", app)
			return next(c)
		}
	})

	// Initialize all the API handlers
	initHTTPHandler(srv)
	srv.Logger.Fatal(srv.Start(fmt.Sprintf(":%d", app.Port)))
	return srv
}