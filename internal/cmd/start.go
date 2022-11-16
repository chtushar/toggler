package cmd

import (
	"fmt"

	"github.com/chtushar/toggler/internal/configs"
	"github.com/chtushar/toggler/internal/logger"
	"github.com/chtushar/toggler/internal/server"
	"github.com/spf13/cobra"
)

func start(cmd *cobra.Command, _ []string) {
	cfg := configs.Get()
	fmt.Println(cfg)
	log := logger.New(&logger.Config{Production: true})

	srv := server.NewServer(&server.Config{
		Port:   cfg.Port,
		Logger: log,
	})

	go srv.Listen()

	srv.WaitForShutdown()
}
