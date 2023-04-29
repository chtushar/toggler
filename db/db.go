package db

import "github.com/jackc/pgx"

type DB struct {
	Conn *pgx.Conn
}
