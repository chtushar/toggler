package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chtushar/toggler/configs"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type App struct {
	port   int
	dbConn *pgx.Conn
	log    *log.Logger
}

var (
	dbConn *pgx.Conn
	cfg    = configs.Get()
)

func initDB() *pgx.Conn {
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

	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	defer conn.Close(context.Background())
	return conn
}

func initHTTPServer(app *App) *echo.Echo {
	fmt.Println("Initializing HTTP Server")
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
	fmt.Println("Initialized HTTP Server")
	srv.Logger.Fatal(srv.Start(fmt.Sprintf(":%d", app.port)))
	return srv
}

// Init is the entry point for the application
func init() {
	// Initialize the database
	dbConn = initDB()

	app := &App{
		port:   cfg.Port,
		dbConn: dbConn,
		log:    log.New(os.Stdout, "toggler: ", log.LstdFlags),
	}

	// Initialize the HTTP server
	initHTTPServer(app)
}
