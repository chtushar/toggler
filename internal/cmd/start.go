package cmd

import (
	"fmt"

	"github.com/chtushar/toggler/internal/config"
	"github.com/chtushar/toggler/internal/logger"
	"github.com/chtushar/toggler/internal/server"
	"github.com/spf13/cobra"
)

func start(cmd *cobra.Command, _ []string) {
	cfg := config.Get()
	fmt.Println(cfg)
	log := logger.New(&logger.Config{Production: true})

	srv := server.NewServer(&server.Config{
		Logger: log,
	})

	go srv.Listen()

	srv.WaitForShutdown()
}
