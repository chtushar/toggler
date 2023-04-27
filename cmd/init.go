package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chtushar/toggler/configs"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type App struct {
	port int
	db   *sqlx.DB
	log  *log.Logger
}

var (
	db  *sqlx.DB
	cfg = configs.Get()
)

func initDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name))

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	return db
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
	db = initDB()

	app := &App{
		port: cfg.Port,
		db:   db,
		log:  log.New(os.Stdout, "toggler: ", log.LstdFlags),
	}

	// Initialize the HTTP server
	initHTTPServer(app)
}
