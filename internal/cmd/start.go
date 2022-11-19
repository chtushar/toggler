package cmd

import (
	"github.com/chtushar/toggler/internal/configs"
	"github.com/chtushar/toggler/internal/db"
	"github.com/chtushar/toggler/internal/logger"
	"github.com/chtushar/toggler/internal/server"
	"github.com/spf13/cobra"
)

func start(cmd *cobra.Command, _ []string) {
	cfg := configs.Get()
	log := logger.New(&logger.Config{Production: true})

	dbConn := db.Get(cfg.DB, log)

	srv := server.NewServer(&server.Config{
		Port:   cfg.Port,
		Logger: log,
	}, dbConn)

	go srv.Listen()

	srv.WaitForShutdown()
}
