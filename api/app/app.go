package app

import (
	"log"

	"github.com/chtushar/toggler/db/queries"
	"github.com/jackc/pgx/v4/pgxpool"
)

type App struct {
	Port   int
	Jwt    string
	DbConn *pgxpool.Pool
	Q      *queries.Queries
	Log    *log.Logger
}