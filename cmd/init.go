package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chtushar/toggler/configs"
	"github.com/chtushar/toggler/db/queries"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type App struct {
	port   int
	jwt    string
	dbConn *pgxpool.Pool
	q      *queries.Queries
	log    *log.Logger
}

var (
	dbConn *pgxpool.Pool
	cfg    = configs.Get()
)

func getPGXConfig() (*pgxpool.Config, error) {
	sslMode := "prefer"

	if cfg.DB.ForceTLS {
		sslMode = "require"
	}
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		sslMode,
	)

	connConfig, err := pgxpool.ParseConfig(connString)

	if err != nil {
		log.Fatalln("Failed to parse config", err)
		return nil, err
	}

	return connConfig, nil
}

func initDB() *pgxpool.Pool {
	pgxConfig, err := getPGXConfig()

	if err != nil {
		log.Fatal("Failed to get database config", err)
		os.Exit(1)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	return pool
}

func initHTTPServer(app *App) *echo.Echo {
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
	initHTTPHandler(srv, app)
	srv.Logger.Fatal(srv.Start(fmt.Sprintf(":%d", app.port)))
	return srv
}
