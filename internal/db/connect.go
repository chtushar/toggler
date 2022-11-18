package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/chtushar/toggler/internal/configs"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DB struct {
	*sqlx.DB
}

var (
	once   sync.Once
	dbConn *DB
)

func Get(ctx context.Context, cfg *configs.DB, logger *zap.Logger) *DB {
	once.Do(func() {
		dbConn = New(ctx, cfg, logger)
		fmt.Println("Connected to database")
	})
	return dbConn
}

func New(ctx context.Context, cfg *configs.DB, logger *zap.Logger) *DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	dbConn, err := sqlx.ConnectContext(ctx, "postgres", dsn)

	if err != nil {
		logger.Panic("Failed to connect to database")
	}

	return &DB{dbConn}
}
