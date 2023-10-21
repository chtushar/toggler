package db

import (
	"context"

	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/db/queries"
)


type dBTransactionFn[T any] func(*queries.Queries) (*T, error)

func WithDBTransaction[T any](app *app.App, ctx context.Context, fn dBTransactionFn[T]) (*T, error) {
	tx, err := app.DbConn.Begin(ctx)
	if err != nil {
		app.Log.Println("Something went wrong", err)
		return nil, err
	}

	defer tx.Rollback(ctx)
	qtx := app.Q.WithTx(tx)

	i, err := fn(qtx)

	tx.Commit(ctx)

	return i, err
}