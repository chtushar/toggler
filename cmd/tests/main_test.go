package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/chtushar/toggler/db/queries"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
)

type App struct {
	port   int
	jwt    string
	dbConn *pgx.Conn
	q      *queries.Queries
	log    *log.Logger
}

var (
	app *App
	pgxConfig *pgx.ConnConfig
)

var (
	dbConn *pgx.Conn
)

func TestMain(m *testing.M) {
	// Start Docker container
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("postgres", "latest", []string{"POSTGRES_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start PostgreSQL container: %s", err)
	}

	// Set up connection parameters
	dbURL := fmt.Sprintf("host=localhost port=%s user=postgres password=secret dbname=postgres sslmode=disable", resource.GetPort("5432/tcp"))
	parsedPgxConfig, err := pgx.ParseConfig(dbURL)
	pgxConfig = parsedPgxConfig

	if err != nil {
		log.Fatalf("Couldn't parse the config")
	}

	// Wait for the container to be ready
	if err := pool.Retry(func() error {
		ctx := context.Background()

		conn, err := pgx.ConnectConfig(ctx, pgxConfig)

		if err != nil {
			return err
		}

		dbConn = conn
		
		defer dbConn.Close(ctx)

		return dbConn.Ping(ctx)
	}); err != nil {
		log.Fatalf("Could not connect to PostgreSQL container: %s", err)
	}

	// Run the tests
	code := m.Run()

	// Clean up Docker container
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge PostgreSQL container: %s", err)
	}

	// Create app
	app = &App {
		port: 9091,
		jwt: "jwt_secret",
		dbConn: dbConn,
		q: queries.New(dbConn),
		log: log.New(os.Stdout, "toggler: ", log.LstdFlags),	
	}

	os.Exit(code)
}
