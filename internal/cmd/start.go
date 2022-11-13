package cmd

import (
	"github.com/chtushar/toggler.in/internal/logger"
	"github.com/chtushar/toggler.in/internal/server"
	"github.com/spf13/cobra"
)

func start(cmd *cobra.Command, _ []string) {
	log := logger.New(&logger.Config{Production: true})

	srv := server.NewServer(&server.Config{
		Logger: log,
	})

	go srv.Listen()

	srv.WaitForShutdown()
}
