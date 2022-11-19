package db

import (
	"fmt"
	"sync"

	"github.com/chtushar/toggler/internal/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type DB struct {
	*sqlx.DB
}

var (
	once   sync.Once
	dbConn *DB
)

func Get(cfg *configs.DB, logger *zap.Logger) *DB {
	once.Do(func() {
		dbConn = New(cfg, logger)
		fmt.Println("Connected to database")
	})
	return dbConn
}

func New(cfg *configs.DB, logger *zap.Logger) *DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Name)
	dbConn, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		logger.Panic("Failed to connect to database")
	}

	return &DB{dbConn}
}
