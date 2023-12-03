package main

import (
	"context"
	"embed"
	"os"
	"os/signal"
	"syscall"

	"github.com/chtushar/toggler/cmd"
	"github.com/chtushar/toggler/configs"
)

//go:embed *.yaml
var configExample embed.FS


func main() {
	// Embed static files
	configs.ConfigExample = configExample

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)
	defer cancel()

	cmd.Execute(ctx)
}
