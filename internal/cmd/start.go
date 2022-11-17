package cmd

import (
	"fmt"

	"github.com/chtushar/toggler/internal/configs"
	"github.com/chtushar/toggler/internal/logger"
	"github.com/chtushar/toggler/internal/server"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

func start(cmd *cobra.Command, _ []string) {
	cfg := configs.Get()
	log := logger.New(&logger.Config{Production: true})

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)

	dbConn, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Panic("Failed to connect to database")
	}

	srv := server.NewServer(&server.Config{
		Port:   cfg.Port,
		Logger: log,
	}, dbConn)

	go srv.Listen()

	srv.WaitForShutdown()
}
