package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chtushar/toggler/configs"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

var (
	dbConn *pgxpool.Pool
	cfg    *configs.Config
)

func getConfigFromFile() {
	cfg = configs.Get()
}

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

// type qLogger struct {
// }

// func (l *qLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
//   if level == pgx.LogLevelInfo && msg == "Query" {
//     fmt.Printf("SQL:\n%s\nARGS:%v\n", data["sql"], data["args"])
//   }
// }


func initDB() *pgxpool.Pool {
	pgxConfig, err := getPGXConfig()

	if err != nil {
		log.Fatal("Failed to get database config", err)
		os.Exit(1)
	}

	// pgxConfig.ConnConfig.Logger = &qLogger{}
	pool, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	return pool
}
